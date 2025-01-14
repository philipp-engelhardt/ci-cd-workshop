package model

import (
	"github.com/google/uuid"
)

type User struct {
	Uuid        uuid.UUID `json:"uuid"`
	DisplayName string    `json:"displayName"`
	Balance     float32   `json:"balance"`
}

type Transaction struct {
	Uuid        uuid.UUID `json:"uuid"`
	ArticleUuid uuid.UUID `json:"article"`
	UserUuid    uuid.UUID `json:"user"`
	Balance     float32   `json:"balance"`
	// TODO: timestamp
}

type Article struct {
	Uuid        uuid.UUID `json:"uuid"`
	Archived    bool      `json:"archived"`
	Available   bool      `json:"available"`
	DisplayName string    `json:"displayName"`
	Price       float32   `json:"price"`
}
