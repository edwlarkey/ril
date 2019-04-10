package bolt

import (
	"flag"
	"log"
	"os"
	"testing"
	"time"

	"github.com/asdine/storm"
	"github.com/edwlarkey/ril/pkg/models"
)

var db *storm.DB
var code int

func TestMain(m *testing.M) {

	flag.Parse()

	if !testing.Short() {
		db, err := storm.Open("tmp.db")
		if err != nil {
			log.Fatal(err)
		}
		article := models.Article{
			Title:     "The Constitution of the United States",
			Content:   "<p>We the People of the United States, in Order to form a more perfect Union, establish Justice, insure domestic Tranquility, provide for the common defence, promote the general Welfare, and secure the Blessings of Liberty to ourselves and our Posterity, do ordain and establish this Constitution for the United States of America.</p>",
			URL:       "https://www.archives.gov/founding-docs/constitution-transcript",
			Created:   time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
			Completed: 0,
		}

		db.Save(&article)

		code = m.Run()

		db.Close()
	} else {
		code = m.Run()
	}

	os.Exit(code)
}
