package inmemory

import (
	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
)

func (imp *InMemoryPersistence) GetAllArticles() ([]model.Article, error) {
	return imp.articles, nil
}
