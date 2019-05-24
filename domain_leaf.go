package idleaf

import (
	"database/sql"
)

type domainLeaf interface {
	Reset(int64, bool) error
	Current() int64
	Gen() (id int64, err error)
}

func newDomainLeaf(db *sql.DB, domain string, table string, idOffset int64) (domainLeaf, error) {
	return newDomainLeafThreadSafe(db, domain, table, idOffset)
}
