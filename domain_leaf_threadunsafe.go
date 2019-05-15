package idleaf

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// using innodb row level lock
const (
	SqlCreateTableIfNotExist = `CREATE TABLE IF NOT EXISTS %s (
    	domain varchar(30) DEFAULT NULL,
    	id bigint(20) NOT NULL,
   		UNIQUE KEY domain (domain)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	SqlFmtSelForUp     = "SELECT id FROM %s where domain='%s' FOR UPDATE"
	SqlFmtAddId        = "UPDATE %s SET id = id + %d where domain='%s'"
	SqlFmtUpId         = "UPDATE %s SET id = %d where domain='%s'"
	SqlFmtInsertDomain = "INSERT INTO %s(domain,id) VALUES('%s',%d)"
	BuffedCount        = 2000
)

type DomainLeafThreadUnsafe struct {
	table       string // leaf table name
	domain      string // domain name
	buffed      int64  // get buffed ids from db once
	idCurrent   int64  // domain's current id
	idBuffedMax int64  // max id buffed
	db          *sql.DB
}

func newDomainLeafThreadUnsafe(db *sql.DB, domain string, leafName string, idOffset int64) (*DomainLeafThreadUnsafe, error) {
	leaf := &DomainLeafThreadUnsafe{db: db, domain: domain, buffed: BuffedCount, table: leafName}
	return leaf, leaf.Reset(idOffset, false)
}

func (p *DomainLeafThreadUnsafe) Reset(idOffset int64, force bool) error {
	var err error
	// create table if not exist
	_, err = p.db.Exec(fmt.Sprintf(SqlCreateTableIfNotExist, p.table))
	if err != nil {
		return err
	}

	// create domain row if not exist
	p.idCurrent, err = p.getIdFromDb(false)
	if err != nil {
		if err == ErrDomainLeafLost {
			sqlInsertDomain := fmt.Sprintf(SqlFmtInsertDomain, p.table, p.domain, idOffset)
			_, err = p.db.Exec(sqlInsertDomain)
			if err != nil {
				return err
			}
			p.idCurrent = idOffset
		}
	} else {
		if force {
			tx, err := p.db.Begin()
			if err != nil {
				return err
			}
			sqlUpdateId := fmt.Sprintf(SqlFmtUpId, p.table, idOffset, p.domain)
			_, err = p.db.Exec(sqlUpdateId)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
			_ = tx.Commit()
		}
	}
	p.idBuffedMax = p.idCurrent
	return err
}

func (p *DomainLeafThreadUnsafe) Init() (err error) {
	if p.idCurrent, err = p.getIdFromDb(false); err != nil {
		return
	}
	p.idBuffedMax = p.idCurrent
	return
}

func (p *DomainLeafThreadUnsafe) Current() int64 {
	return p.idCurrent
}

func (p *DomainLeafThreadUnsafe) getIdFromDb(buffed bool) (id int64, err error) {

	sqlSelForUp := fmt.Sprintf(SqlFmtSelForUp, p.table, p.domain)

	// begin transaction
	tx, err := p.db.Begin()
	if err != nil {
		return
	}

	rows, err := tx.Query(sqlSelForUp)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	defer rows.Close()

	found := false
	// must clear query result
	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			_ = tx.Rollback()
			return
		}
		found = true
	}

	if found == false {
		err = ErrDomainLeafLost
		_ = tx.Rollback()
		return
	}

	if buffed {
		if _, err = tx.Exec(fmt.Sprintf(SqlFmtAddId, p.table, p.buffed, p.domain)); err != nil {
			_ = tx.Rollback()
			return
		}
	}
	_ = tx.Commit()

	return id, nil
}

func (p *DomainLeafThreadUnsafe) Gen() (id int64, err error) {
	if p.idBuffedMax < p.idCurrent+1 {
		if id, err = p.getIdFromDb(true); err != nil {
			return
		}
		p.idCurrent = id
		p.idBuffedMax = p.idCurrent + p.buffed
	}

	p.idCurrent++
	return p.idCurrent, nil
}
