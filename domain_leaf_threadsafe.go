package idleaf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type DomainLeafThreadsafe struct {
	leaf *DomainLeafThreadUnsafe
	sync.Mutex
}

func newDomainLeafThreadSafe(db *sql.DB, domain string, table string, idOffset int64) (*DomainLeafThreadsafe, error) {
	leaf, err := newDomainLeafThreadUnsafe(db, domain, table, idOffset)
	if err != nil {
		return nil, err
	}
	return &DomainLeafThreadsafe{leaf: leaf}, nil
}

func (p *DomainLeafThreadsafe) Reset(idOffset int64, force bool) error {
	p.Lock()
	defer p.Unlock()
	return p.leaf.Reset(idOffset, force)
}

func (p *DomainLeafThreadsafe) GetDomain() string {
	return p.leaf.domain
}

func (p *DomainLeafThreadsafe) Current() int64 {
	p.Lock()
	defer p.Unlock()
	return p.leaf.Current()
}

func (p *DomainLeafThreadsafe) Gen() (int64, error) {
	p.Lock()
	defer p.Unlock()
	return p.leaf.Gen()
}
