package idleaf

import (
	"database/sql"
	"fmt"
	"sync"
)

type IdLeaf struct {
	option  *Option
	db      *sql.DB
	syncMap sync.Map
}

func NewIdLeaf(option *Option) (p *IdLeaf, err error) {
	p = new(IdLeaf)
	p.option = option

	//root:@tcp(127.0.0.1:3306)/test?charset=utf8
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		p.option.DbUser,
		p.option.DbPass,
		p.option.DbHost,
		p.option.DbPort,
		p.option.DbName,
	)
	if p.db, err = sql.Open(p.option.DbProto, url); err != nil {
		return
	}
	err = p.db.Ping()
	return
}

func (p *IdLeaf) GenId(domain string) (int64, error) {
	var leaf DomainLeafIF
	var err error
	if lif, ok := p.syncMap.Load(domain); !ok {
		leaf, err = NewDomainLeaf(p.db, domain, p.option.LeafTable, p.option.IdOffset)
		if err != nil {
			return 0, err
		}
		p.syncMap.Store(domain, leaf)
	} else {
		leaf, _ = lif.(DomainLeafIF)
	}
	return leaf.Gen()
}

var idLeaf *IdLeaf = nil

func Init(option *Option) (err error) {
	idLeaf, err = NewIdLeaf(option)
	return
}
