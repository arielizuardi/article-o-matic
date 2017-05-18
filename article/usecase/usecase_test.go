package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/arielizuardi/article-o-matic/article"
	"github.com/arielizuardi/article-o-matic/article/repository/mocks"
	"github.com/arielizuardi/article-o-matic/article/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStore(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("FindByURL", mock.AnythingOfType("string")).Return(nil, nil)
	repo.On("Store", mock.AnythingOfType("*article.Article")).Return(nil)

	u := usecase.NewArticleUsecase(repo)

	a := &article.Article{
		Title:     "this is title",
		Content:   "this is content",
		URL:       "http://example.com/1",
		CreatedAt: time.Now(),
	}

	err := u.Store(a)
	assert.NoError(t, err)

	repo.AssertCalled(t, "FindByURL", mock.AnythingOfType("string"))
	repo.AssertCalled(t, "Store", mock.AnythingOfType("*article.Article"))
}

func TestStoreAndArticleAlreadyExists(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("FindByURL", mock.AnythingOfType("string")).Return(&article.Article{}, nil)
	repo.On("Store", mock.AnythingOfType("*article.Article")).Return(nil)

	u := usecase.NewArticleUsecase(repo)

	a := &article.Article{
		Title:     "this is title",
		Content:   "this is content",
		URL:       "http://example.com/1",
		CreatedAt: time.Now(),
	}

	err := u.Store(a)
	assert.Error(t, err)

	repo.AssertCalled(t, "FindByURL", mock.AnythingOfType("string"))
	repo.AssertNotCalled(t, "Store", mock.AnythingOfType("*article.Article"))
}

func TestStoreAndRepositoryReturnError(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("FindByURL", mock.AnythingOfType("string")).Return(nil, errors.New(`Whoops!`))
	repo.On("Store", mock.AnythingOfType("*article.Article")).Return(nil)

	u := usecase.NewArticleUsecase(repo)

	a := &article.Article{
		Title:     "this is title",
		Content:   "this is content",
		URL:       "http://example.com/1",
		CreatedAt: time.Now(),
	}

	err := u.Store(a)
	assert.Error(t, err)

	repo.AssertCalled(t, "FindByURL", mock.AnythingOfType("string"))
	repo.AssertNotCalled(t, "Store", mock.AnythingOfType("*article.Article"))
}
