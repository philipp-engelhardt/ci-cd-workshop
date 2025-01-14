package inmemory

import (
	"github.com/google/uuid"
	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
	"github.com/rubenhoenle/ci-cd-lecture-project-template/util"
)

func (imp *InMemoryPersistence) GetAllUsers() ([]model.User, error) {
	return imp.users, nil
}

func (imp *InMemoryPersistence) CreateUser(user model.User) error {
	imp.users = append(imp.users, user)
	return nil
}

func (imp *InMemoryPersistence) GetUser(uuid uuid.UUID) (*model.User, error) {
	for _, u := range imp.users {
		if u.Uuid == uuid {
			return &u, nil
		}
	}
	return nil, nil
}

func (imp *InMemoryPersistence) UpdateUser(user model.User) error {
	for idx, u := range imp.users {
		if u.Uuid == user.Uuid {
			imp.users[idx] = user
			break
		}
	}
	return nil
}

func (imp *InMemoryPersistence) DeleteUser(uuid uuid.UUID) error {
	for idx, u := range imp.users {
		if u.Uuid == uuid {
			imp.users = util.RemoveUser(imp.users, idx)
			break
		}
	}
	return nil
}
