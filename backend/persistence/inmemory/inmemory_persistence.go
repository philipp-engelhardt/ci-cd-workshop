package inmemory

import (
	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
)

// implements the persistence interface
type InMemoryPersistence struct {
	users        []model.User
	articles     []model.Article
	transactions []model.Transaction
}

func NewInMemoryPersistence() *InMemoryPersistence {
	return &InMemoryPersistence{users: []model.User{}, articles: []model.Article{}, transactions: []model.Transaction{}}
}
