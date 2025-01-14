package inmemory

import (
	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
)

func (imp *InMemoryPersistence) GetAllTransactions() ([]model.Transaction, error) {
	return imp.transactions, nil
}
