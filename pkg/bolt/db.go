package bolt

import (
	"github.com/asdine/storm"
)

type DB struct {
	DB *storm.DB
}

func (d *DB) Connect(dsn string) error {
	db, err := storm.Open(dsn)

	if err != nil {
		return err
	}

	d.DB = db
	return nil
}

func (d *DB) Close() {
	d.DB.Close()
}
