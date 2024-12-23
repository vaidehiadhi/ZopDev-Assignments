package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/vaidehiadhi/threeLayerArc/models"
	"gotest.tools/v3/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockService struct {
}

func (m mockService) GetUser(name string) (*models.User, error) {
	if name == "validUser" {
		return &models.User{Name: "vai vai", Age: 20, Phone: 823832789, Email: "vai@c.com"}, nil
	}
	return nil, errors.New("user not found")
}

func (m mockService) AddUser(user *models.User) error {
	if user.Name == "duplicateUser" {
		return errors.New("user already exists")
	}
	return nil
}

func (m mockService) UpdateUser(name string, user *models.User) error {
	if name != "validUser" {
		return errors.New("user not found")
	}
	return nil
}

func (m mockService) DeleteUser(name string) error {
	if name != "validUser" {
		return errors.New("user not found")
	}
	return nil
}

func TestUserHandler_GetUser(t *testing.T) {
	handler := NewUserHandler(mockService{})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/validUser", nil)
		req = mux.SetURLVars(req, map[string]string{"name": "validUser"})
		rec := httptest.NewRecorder()
		handler.GetUser(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)
		expectedBody := `{"Name":"vai vai","Age":20,"Phone":823832789,"Email":"vai@c.com"}`
		assert.Equal(t, rec.Body.String(), expectedBody)
	})

	t.Run("failure", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/invalidUser", nil)
		req = mux.SetURLVars(req, map[string]string{"name": "invalidUser"})
		rec := httptest.NewRecorder()
		handler.GetUser(rec, req)
		assert.Equal(t, rec.Code, http.StatusInternalServerError)
	})
}

func TestUserHandler_AddUser(t *testing.T) {
	handler := NewUserHandler(mockService{})

	t.Run("success", func(t *testing.T) {
		user := models.User{Name: "newUser", Age: 30, Email: "new@example.com", Phone: 9876543210}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		rec := httptest.NewRecorder()
		handler.AddUser(rec, req)
		assert.Equal(t, rec.Code, http.StatusBadRequest)

		expectedBody := `{"name":"newUser","age":30,"phone":9876543210,"email":"new@example.com"}`
		assert.Equal(t, rec.Body.String(), expectedBody)
	})

	t.Run("failure - duplicate user", func(t *testing.T) {
		user := models.User{Name: "duplicateUser", Age: 30, Email: "duplicate@example.com", Phone: 9876543210}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		rec := httptest.NewRecorder()
		handler.AddUser(rec, req)
		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("failure - invalid body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte("{invalid json}")))
		rec := httptest.NewRecorder()
		handler.AddUser(rec, req)
		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})
}

func TestUserHandler_UpdateUser(t *testing.T) {
	handler := NewUserHandler(mockService{})

	t.Run("success", func(t *testing.T) {
		user := models.User{Name: "updatedUser", Age: 35, Email: "updated@example.com", Phone: 1231231234}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPut, "/users/validUser", bytes.NewBuffer(body))
		req = mux.SetURLVars(req, map[string]string{"name": "validUser"})
		rec := httptest.NewRecorder()
		handler.UpdateUser(rec, req)
		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("failure - user not found", func(t *testing.T) {
		user := models.User{Name: "unknownUser", Age: 35, Email: "unknown@example.com", Phone: 1231231234}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPut, "/users/invalidUser", bytes.NewBuffer(body))
		req = mux.SetURLVars(req, map[string]string{"name": "invalidUser"})
		rec := httptest.NewRecorder()
		handler.UpdateUser(rec, req)
		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("failure - invalid body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/users/validUser", bytes.NewBuffer([]byte("{invalid json}")))
		req = mux.SetURLVars(req, map[string]string{"name": "validUser"})
		rec := httptest.NewRecorder()
		handler.UpdateUser(rec, req)
		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})
}

func TestUserHandler_DeleteUser(t *testing.T) {
	handler := NewUserHandler(mockService{})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/validUser", nil)
		req = mux.SetURLVars(req, map[string]string{"name": "validUser"})
		rec := httptest.NewRecorder()
		handler.DeleteUser(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)
	})

	t.Run("failure - user not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/invalidUser", nil)
		req = mux.SetURLVars(req, map[string]string{"name": "invalidUser"})
		rec := httptest.NewRecorder()
		handler.DeleteUser(rec, req)
		assert.Equal(t, rec.Code, http.StatusInternalServerError)
	})
}
