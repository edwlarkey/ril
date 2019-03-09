package mysql

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest"
)

var db *sql.DB
var code int

func TestMain(m *testing.M) {

	flag.Parse()

	if !testing.Short() {
		pool, err := dockertest.NewPool("")
		if err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}
		// pulls an image, creates a container based on it and runs it
		resource, err := pool.Run("mariadb", "latest", []string{"MYSQL_ROOT_PASSWORD=secret", "MYSQL_DATABASE=test"})
		if err != nil {
			log.Fatalf("Could not start resource: %s", err)
		}

		// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
		if err := pool.Retry(func() error {
			var err error
			db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/test?parseTime=true&multiStatements=true", resource.GetPort("3306/tcp")))
			if err != nil {
				return err
			}
			return db.Ping()
		}); err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}

		// Read the setup SQL script from file and execute the statements.
		script, err := ioutil.ReadFile("./testdata/setup.sql")
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			log.Fatal(err)
		}

		code = m.Run()

		script, err = ioutil.ReadFile("./testdata/teardown.sql")
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			log.Fatal(err)
		}

		db.Close()

		// You can't defer this because os.Exit doesn't care for defer
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	} else {
		code = m.Run()
	}

	os.Exit(code)
}
