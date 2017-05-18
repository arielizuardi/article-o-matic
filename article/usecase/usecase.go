package usecase

import (
	"github.com/arielizuardi/article-o-matic/article"
	"github.com/arielizuardi/article-o-matic/article/repository"
)

// ArticleUsecase defines article usecase
type ArticleUsecase interface {
	Store(a *article.Article) error
	FindByID(id int) (*article.Article, error)
	Update(id int, a *article.Article) error
	Delete(id int) error
}

type articleUsecase struct {
	ArticleRepository repository.ArticleRepository
}

func (u *articleUsecase) Store(a *article.Article) error {
	articleExists, err := u.ArticleRepository.FindByURL(a.URL)
	if err != nil {
		return err
	}

	if articleExists != nil {
		return article.ErrAlreadyExists
	}

	return u.ArticleRepository.Store(a)
}

func (u *articleUsecase) FindByID(id int) (*article.Article, error) {
	return nil, nil
}

func (u *articleUsecase) Update(id int, a *article.Article) error {
	return nil
}

func (u *articleUsecase) Delete(id int) error {
	return nil
}

func NewArticleUsecase(r repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{r}
}
