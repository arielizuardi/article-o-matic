package repository

import "github.com/arielizuardi/article-o-matic/article"

type ArticleRepository interface {
	Store(a *article.Article) error
	FindByURL(url string) (*article.Article, error)
}
