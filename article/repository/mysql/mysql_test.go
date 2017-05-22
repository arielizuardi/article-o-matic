package mysql_test

import (
	"database/sql"
	"testing"

	"github.com/arielizuardi/article-o-matic/article"
	"github.com/arielizuardi/article-o-matic/article/repository/mysql"
	"github.com/arielizuardi/article-o-matic/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// MySQLTest ...
type MySQLTest struct {
	db.MySQLSuite
	conn *sql.DB
}

func TestMySQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip placement repository test")
	}

	suite.Run(t, new(MySQLTest))
}

func (s *MySQLTest) SetupTest() {
	errs, ok := db.RunAllMigrations(s.DSN)
	assert.True(s.T(), ok)
	assert.Len(s.T(), errs, 0)
}

func (s *MySQLTest) TearDownTest() {
}

func (s *MySQLTest) TestStore() {
	repo := mysql.NewMySQLArticleRepository(s.DBConn)
	a := &article.Article{Title: `this is title`, Content: `this is content`, URL: `http://example.com`}
	err := repo.Store(a)
	assert.NoError(s.T(), err)

	var ct int
	query := `SELECT count(*) as ct FROM article where URL = ? `
	result := s.DBConn.QueryRow(query, a.URL)
	err = result.Scan(&ct)
	assert.NoError(s.T(), err)
	assert.NotZero(s.T(), ct)
}
