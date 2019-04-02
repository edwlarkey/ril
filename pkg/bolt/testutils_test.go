package bolt

import (
	"flag"
	"log"
	"os"
	"testing"

	"github.com/asdine/storm"
)

var db *storm.DB
var code int

func TestMain(m *testing.M) {

	flag.Parse()

	if !testing.Short() {
		db = storm.Open("tmp.db")
		if err != nil {
			log.Fatal(err)
		}

		code = m.Run()

		db.Close()
	} else {
		code = m.Run()
	}

	os.Exit(code)
}
