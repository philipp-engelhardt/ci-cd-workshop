package util

import (
	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
)

/* remove the given index from the slice */
func RemoveUser(s []model.User, i int) []model.User {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func RemoveArticle(s []model.Article, i int) []model.Article {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func RemoveTransaction(s []model.Transaction, i int) []model.Transaction {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
