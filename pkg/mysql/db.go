package mysql

import "database/sql"

type DB struct {
	DB *sql.DB
}

func (d *DB) Connect(dsn string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}

	d.DB = db
	return nil
}

func (d *DB) Close() {
	d.DB.Close()
}
