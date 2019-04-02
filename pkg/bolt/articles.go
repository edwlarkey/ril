package bolt

import (
	"time"

	"github.com/asdine/storm"
	"github.com/edwlarkey/ril/pkg/models"
)

// Insert adds an article to the DB
func (m *DB) InsertArticle(title, content, url string) (int, error) {
	article := models.Article{
		Title:     title,
		Content:   content,
		URL:       url,
		Created:   time.Now(),
		Completed: 0,
	}

	err := m.DB.Save(&article)

	if err != nil {
		return 0, err
	}

	return article.ID, nil

}

// Get gets a single article from the DB
func (m *DB) GetArticle(id int) (*models.Article, error) {
	a := models.Article{}

	err := m.DB.One("ID", id, &a)
	if err == storm.ErrNotFound {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return &a, nil

}

// Latest gets the 10 most recent articles from the DB
func (m *DB) LatestArticles() ([]*models.Article, error) {
	a := []models.Article{}
	articles := []*models.Article{}

	err := m.DB.All(&a, storm.Limit(10))

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(a); i++ {
		articles = append(articles, &a[i])
	}

	return articles, nil
}
