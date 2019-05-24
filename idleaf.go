package idleaf

import (
	"database/sql"
	"fmt"
	"sync"
)

type leaf struct {
	option  *Option
	db      *sql.DB
	syncMap sync.Map
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
	var err error
	if lif, ok := p.syncMap.Load(domain); !ok {
		leaf, err = newDomainLeaf(p.db, domain, p.option.LeafTable, p.option.IdOffset)
		if err != nil {
			return 0, err
		}
		p.syncMap.Store(domain, leaf)
	} else {
		leaf, _ = lif.(domainLeaf)
	}
	return leaf.Gen()
}

var idLeaf *leaf = nil

func Init(option *Option) (err error) {
	idLeaf, err = newLeaf(option)
	return
}
