package idleaf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DomainLeaf interface {
	Reset(int64, bool) error
	Current() int64
	Gen() (id int64, err error)
}

func NewDomainLeaf(db *sql.DB, domain string, table string, idOffset int64) (DomainLeaf, error) {
	return newDomainLeafThreadSafe(db, domain, table, idOffset)
}
