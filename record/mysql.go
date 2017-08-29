package record

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nicholaskh/cottage/config"
)

type Mysql struct {
	config *config.RecordConfig
	conn   *sql.DB
}

func newMysql(config *config.RecordConfig) *Mysql {
	this := new(Mysql)
	this.config = config
	return this
}

func (this *Mysql) open() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", this.config.Username, this.config.Password, this.config.Host, this.config.Port, this.config.Db)
	db, err := sql.Open("mysql", dsn)
	return db, err
}

func (this *Mysql) Record(ver, name, url string) (err error) {
	db, err := this.open()
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from package where name=? and version=?", name, ver)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		return errors.New("package already exists")
	}

	stmtIns, err := db.Prepare("INSERT INTO package(name, version, url) VALUES(?, ?, ?)")
	if err != nil {
		return
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(name, ver, url)
	if err != nil {
		return
	}
	return
}

func (this *Mysql) Search(word string) (names []string, err error) {
	db, err := this.open()
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("select distinct name from package where name like '%%%s%%'", word))
	if err != nil {
		return
	}
	defer rows.Close()

	names = make([]string, 0)
	var name string
	for rows.Next() {
		rows.Scan(&name)
		names = append(names, name)
	}
	return
}

func (this *Mysql) MaxVersion(name string) (version string, err error) {
	db, err := this.open()
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query("select max(version) from package where name=?", name)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&version)
	}
	return
}

func (this *Mysql) GetUrl(name, version string) (url string, err error) {
	db, err := this.open()
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query("select url from package where name=? and version=?", name, version)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&url)
	}
	return
}
