package mysql

import (
	"database/sql"
	"time"

	"github.com/arielizuardi/article-o-matic/article"
)

type MySQLArticleRepsository struct {
	Conn *sql.DB
}

func (m *MySQLArticleRepsository) Store(a *article.Article) error {
	insertStmt := `INSERT INTO article (title, content, url, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	stmt, err := m.Conn.Prepare(insertStmt)
	if err != nil {
		return err
	}

	defer stmt.Close()

	now := time.Now()
	r, err := stmt.Exec(a.Title, a.Content, a.URL, now, now)

	if err != nil {
		return err
	}

	var id int64
	id, err = r.LastInsertId()
	if err != nil {
		return err
	}

	a.ID = int(id)

	return nil
}

func (m *MySQLArticleRepsository) FindByURL(url string) (*article.Article, error) {
	return nil, nil
}

func NewMySQLArticleRepository(conn *sql.DB) *MySQLArticleRepsository {
	return &MySQLArticleRepsository{Conn: conn}
}
