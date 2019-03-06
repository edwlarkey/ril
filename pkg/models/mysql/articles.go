package mysql

import (
	"database/sql"

	"github.com/edwlarkey/ril/pkg/models"
)

type ArticleModel struct {
	DB *sql.DB
}

// Insert adds an article to the DB
func (m *ArticleModel) Insert(title, content, url string) (int, error) {
	stmt := `INSERT INTO articles (title, content, url, created, completed)
    VALUES(?, ?, ?, UTC_TIMESTAMP(), 0)`

	result, err := m.DB.Exec(stmt, title, content, url)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get gets a single article from the DB
func (m *ArticleModel) Get(id int) (*models.Article, error) {
	stmt := `SELECT id, title, content, url, created, completed FROM articles
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	a := &models.Article{}

	err := row.Scan(&a.ID, &a.Title, &a.Content, &a.URL, &a.Created, &a.Completed)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return a, nil
}

// Latest gets the 10 most recent articles from the DB
func (m *ArticleModel) Latest() ([]*models.Article, error) {
	stmt := `SELECT id, title, content, url, created, completed FROM articles
    WHERE completed != 1 ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	articles := []*models.Article{}

	for rows.Next() {
		a := &models.Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.URL, &a.Created, &a.Completed)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
