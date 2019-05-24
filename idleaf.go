package idleaf

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/timestee/golib/singleflight"
)

type leaf struct {
	option  *Option
	db      *sql.DB
	syncMap sync.Map
	sf      singleflight.Group
}

func newLeaf(option *Option) (p *leaf, err error) {
	p = &leaf{option: option}

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		p.option.DbUser,
		p.option.DbPass,
		p.option.DbHost,
		p.option.DbPort,
		p.option.DbName,
	)
	fmt.Println(url)

	if p.db, err = sql.Open(p.option.DbProto, url); err != nil {
		return
	}
	err = p.db.Ping()
	return
}

func (p *leaf) genId(domain string) (int64, error) {
	var leaf domainLeaf

	if lif, ok := p.syncMap.Load(domain); ok {
		leaf, _ = lif.(domainLeaf)
	} else {
		lif, err := p.sf.Do(domain, func() (i interface{}, e error) {
			// newDomainLeaf will check table, create domain row if not exist
			leaf, err := newDomainLeaf(p.db, domain, p.option.LeafTable, p.option.IdOffset)
			if err == nil {
				p.syncMap.Store(domain, leaf)
			}
			return leaf, err
		})
		if err != nil {
			return 0, err
		}
		leaf, _ = lif.(domainLeaf)
	}
	return leaf.Gen()
}

var idLeaf *leaf

// InitLeaf init the global leaf with given option
func InitLeaf(option *Option) (err error) {
	idLeaf, err = newLeaf(option)
	return
}
