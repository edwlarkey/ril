package mock

import (
	"time"

	"github.com/edwlarkey/ril/pkg/models"
)

var mockArticle = &models.Article{
	ID:        1,
	Title:     "The Constitution of the United States",
	Content:   "<p>We the People of the United States, in Order to form a more perfect Union, establish Justice, insure domestic Tranquility, provide for the common defence, promote the general Welfare, and secure the Blessings of Liberty to ourselves and our Posterity, do ordain and establish this Constitution for the United States of America.</p>",
	URL:       "https://www.archives.gov/founding-docs/constitution-transcript",
	Created:   time.Now(),
	Completed: 0,
}

type ArticleModel struct{}

func (m *ArticleModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

func (m *ArticleModel) Get(id int) (*models.Article, error) {
	switch id {
	case 1:
		return mockArticle, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *ArticleModel) Latest() ([]*models.Article, error) {
	return []*models.Article{mockArticle}, nil
}
