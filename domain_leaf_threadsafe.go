package idleaf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type DomainLeafThreadsafe struct {
	leaf *DomainLeafThreadUnsafe
	lock sync.Mutex
}

func newDomainLeafThreadSafe(db *sql.DB, domain string, table string, id_offset int64) (*DomainLeafThreadsafe, error) {
	leaf, err := newDomainLeafThreadUnsafe(db, domain, table, id_offset)
	if err != nil {
		return nil, err
	}
	return &DomainLeafThreadsafe{leaf: leaf}, nil
}

func (p *DomainLeafThreadsafe) Reset(idOffset int64, force bool) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.leaf.Reset(idOffset, force)
}

func (p *DomainLeafThreadsafe) GetDomain() string {
	return p.leaf.domain
}

func (p *DomainLeafThreadsafe) Current() int64 {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.leaf.Current()
}

func (p *DomainLeafThreadsafe) Gen() (int64, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.leaf.Gen()
}
