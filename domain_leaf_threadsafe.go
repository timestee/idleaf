package idleaf

import (
	"database/sql"
	"sync"
)

type domainLeafThreadsafe struct {
	leaf *domainLeafThreadUnsafe
	sync.Mutex
}

func newDomainLeafThreadSafe(db *sql.DB, domain string, table string, idOffset int64) (*domainLeafThreadsafe, error) {
	leaf, err := newDomainLeafThreadUnsafe(db, domain, table, idOffset)
	if err != nil {
		return nil, err
	}
	return &domainLeafThreadsafe{leaf: leaf}, nil
}

func (p *domainLeafThreadsafe) Reset(idOffset int64, force bool) error {
	p.Lock()
	defer p.Unlock()
	return p.leaf.Reset(idOffset, force)
}

func (p *domainLeafThreadsafe) GetDomain() string {
	return p.leaf.domain
}

func (p *domainLeafThreadsafe) Current() int64 {
	p.Lock()
	defer p.Unlock()
	return p.leaf.Current()
}

func (p *domainLeafThreadsafe) Gen() (int64, error) {
	p.Lock()
	defer p.Unlock()
	return p.leaf.Gen()
}
