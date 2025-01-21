package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
)

type ArticlesTest struct {
}

func (test *ArticlesTest) GetTestName() string {
	return "API - Get all articles"
}

func (test *ArticlesTest) RunTest(baseUrl string) bool {
	resp, err := http.Get(baseUrl + "/api/article")
	if err != nil {
		fmt.Printf("Error: Failed to send request to /api/article: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Expected status 200, got %d\n", resp.StatusCode)
		return false
	}

	var articles []model.Article
	err = json.NewDecoder(resp.Body).Decode(&articles)
	if err != nil {
		fmt.Printf("Error: Failed to decode JSON response: %v\n", err)
		return false
	}

	if len(articles) != 0 {
		fmt.Printf("Error: Wrong number of articles present, expected 0, got %d", len(articles))
		return false
	}

	return true
}
