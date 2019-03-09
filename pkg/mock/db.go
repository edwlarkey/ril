package mock

type DB struct{}

func (d *DB) Connect(dsn string) error {
	return nil
}

func (d *DB) Close() {
	return
}
