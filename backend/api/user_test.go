package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
	"github.com/rubenhoenle/ci-cd-lecture-project-template/persistence/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	persistence := inmemory.NewInMemoryPersistence()
	router := NewRouter(persistence)

	// assert there are some users before
	createdUsers := []model.User{
		{Uuid: uuid.New(), DisplayName: "Luke Skywalker", Balance: 99.9},
		{Uuid: uuid.New(), DisplayName: "Yoda", Balance: -23.0},
	}
	for _, u := range createdUsers {
		err := persistence.CreateUser(u)
		assert.NoError(t, err)
	}

	// fetch all users via the API
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/user", nil)
	router.ServeHTTP(w, req)

	// make sure the API responds correctly
	assert.Equal(t, http.StatusOK, w.Code)
	// check the response body
	var actual []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &actual)
	assert.NoError(t, err)
	expected := []map[string]interface{}{
		{"uuid": createdUsers[0].Uuid.String(), "displayName": "Luke Skywalker", "balance": 99.9},
		{"uuid": createdUsers[1].Uuid.String(), "displayName": "Yoda", "balance": -23.0},
	}
	assert.Equal(t, expected, actual)
}

func TestCreateUser(t *testing.T) {
	persistence := inmemory.NewInMemoryPersistence()
	router := NewRouter(persistence)

	// assert there are no users present before
	users, err := persistence.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(users))

	// create the json body for the request
	user := model.User{DisplayName: "Luke Skywalker", Balance: 99.9}
	userJson, err := json.Marshal(user)
	assert.NoError(t, err)

	// create a user via the API
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/user", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	// assert a single user is present afterwards
	users, err = persistence.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))

	// make sure the API responds correctly
	assert.Equal(t, http.StatusCreated, w.Code)
	// check the response body
	var actual map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &actual)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"uuid": users[0].Uuid.String(), "displayName": "Luke Skywalker", "balance": 99.9}, actual)
}

func TestDeleteUser(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		persistence := inmemory.NewInMemoryPersistence()
		router := NewRouter(persistence)

		// assert there are no users before
		users, err := persistence.GetAllUsers()
		assert.NoError(t, err)
		assert.Equal(t, 0, len(users))

		// create a new user
		user := model.User{Uuid: uuid.New(), DisplayName: "Luke Skywalker", Balance: 99.9}
		err = persistence.CreateUser(user)
		assert.NoError(t, err)
		users, err = persistence.GetAllUsers()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(users))

		// delete the user via the API
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/user/%s", user.Uuid.String()), nil)
		router.ServeHTTP(w, req)

		// make sure the API responds correctly
		assert.Equal(t, http.StatusOK, w.Code)
		// check the response body
		var actual map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &actual)
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"message": "OK"}, actual)

		// assert there are no users again afterwards
		users, err = persistence.GetAllUsers()
		assert.NoError(t, err)
		assert.Equal(t, 0, len(users))
	})

	t.Run("try to delete non-existent user", func(t *testing.T) {
		persistence := inmemory.NewInMemoryPersistence()
		router := NewRouter(persistence)

		// assert there are some users present before
		user := model.User{Uuid: uuid.New(), DisplayName: "Luke Skywalker", Balance: 99.9}
		err := persistence.CreateUser(user)
		assert.NoError(t, err)
		users, err := persistence.GetAllUsers()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(users))

		// delete some non-existent user via the API
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/user/%s", uuid.New().String()), nil)
		router.ServeHTTP(w, req)

		// make sure the API responds correctly
		assert.Equal(t, http.StatusNotFound, w.Code)
		// check the response body
		var actual map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &actual)
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"message": "Not found"}, actual)

		// assert the other users are still there afterwards
		users, err = persistence.GetAllUsers()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(users))
	})

	t.Run("try to delete non-existent user", func(t *testing.T) {
		persistence := inmemory.NewInMemoryPersistence()
		router := NewRouter(persistence)

		// assert there are some users present before
		user := model.User{Uuid: uuid.New(), DisplayName: "Luke Skywalker", Balance: 99.9}
		err := persistence.CreateUser(user)
		assert.NoError(t, err)
		users, err := persistence.GetAllUsers()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(users))

		// try to delete a user using an invalid uuid string (FYI: the uuid below is missing one character)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/api/user/ae1a2d1e-91e4-4a8a-bef9-14e3d118827", nil)
		router.ServeHTTP(w, req)

		// make sure the API responds correctly
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// check the response body
		var actual map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &actual)
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"message": "Invalid uuid"}, actual)

		// assert the other users are still there afterwards
		users, err = persistence.GetAllUsers()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(users))
	})
}
