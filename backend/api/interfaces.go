package api

import (
	"github.com/google/uuid"
	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
)

type Persistence interface {
	GetAllArticles() ([]model.Article, error)
	//GetArticle(uuid.UUID) (*model.Article, error)
	//CreateArticle(model.Article) error
	//UpdateArticle(model.Article) error
	//ArchiveArticle(uuid.UUID) error

	GetAllUsers() ([]model.User, error)
	GetUser(uuid.UUID) (*model.User, error)
	CreateUser(model.User) error
	UpdateUser(model.User) error
	DeleteUser(uuid.UUID) error

	GetAllTransactions() ([]model.Transaction, error)
}
