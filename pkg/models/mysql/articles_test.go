package mysql

import (
	"reflect"
	"testing"
	"time"

	"github.com/edwlarkey/ril/pkg/models"
)

func TestArticlesModelGet(t *testing.T) {
	// Skip the test if the `-short` flag is provided when running the test.
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	// Set up a suite of table-driven tests and expected results.
	tests := []struct {
		name        string
		ID          int
		wantArticle *models.Article
		wantError   error
	}{
		{
			name: "Valid ID",
			ID:   1,
			wantArticle: &models.Article{
				ID:        1,
				Title:     "The Constitution of the United States",
				Content:   "<p>We the People of the United States, in Order to form a more perfect Union, establish Justice, insure domestic Tranquility, provide for the common defence, promote the general Welfare, and secure the Blessings of Liberty to ourselves and our Posterity, do ordain and establish this Constitution for the United States of America.</p>",
				URL:       "https://www.archives.gov/founding-docs/constitution-transcript",
				Created:   time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
				Completed: 0,
			},
			wantError: nil,
		},
		{
			name:        "Zero ID",
			ID:          0,
			wantArticle: nil,
			wantError:   models.ErrNoRecord,
		},
		{
			name:        "Non-existent ID",
			ID:          2,
			wantArticle: nil,
			wantError:   models.ErrNoRecord,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize a connection pool to our test database, and defer a
			// call to the teardown function, so it is always run immediately
			// before this sub-test returns.

			// Create a new instance of the ArticleModel.
			m := ArticleModel{db}

			// Call the ArticleModel.Get() method and check that the return value
			// and error match the expected values for the sub-test.
			article, err := m.Get(tt.ID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if !reflect.DeepEqual(article, tt.wantArticle) {
				t.Errorf("want %v; got %v", tt.wantArticle, article)
			}
		})
	}
}
