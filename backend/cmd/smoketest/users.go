package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
)

type UsersTest struct {
}

func (test *UsersTest) GetTestName() string {
	return "API - Get all users"
}

func (test *UsersTest) RunTest(baseUrl string) bool {
	resp, err := http.Get(baseUrl + "/api/user")
	if err != nil {
		fmt.Printf("Error: Failed to send request to /api/user: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Expected status 200, got %d\n", resp.StatusCode)
		return false
	}

	var users []model.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		fmt.Printf("Error: Failed to decode JSON response: %v\n", err)
		return false
	}

	if len(users) != 0 {
		fmt.Printf("Error: Wrong number of users present, expected 0, got %d", len(users))
		return false
	}

	return true
}
