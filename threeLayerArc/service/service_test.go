package service

import (
	"errors"
	"github.com/vaidehiadhi/threeLayerArc/models"
	"testing"

	"gotest.tools/v3/assert"
)

type mockStore struct {
}

func (m mockStore) GetUser(name string) (*models.User, error) {
	if name == "vaidehi" {
		return &models.User{Name: "vaidehi", Age: 20, Phone: 823832789, Email: "vai@c.com"}, nil
	}
	return nil, errors.New("user not found")
}

func (m mockStore) AddUser(user *models.User) error {
	if user.Name == "duplicateUser" {
		return errors.New("user already exists")
	}
	return nil
}

func (m mockStore) UpdateUser(name string, user *models.User) error {
	if name == "existingUser" {
		return nil
	}
	return errors.New("user not found")
}

func (m mockStore) DeleteUser(name string) error {
	if name == "existingUser" {
		return nil
	}
	return errors.New("user not found")
}

func TestUserService_GetUser(t *testing.T) {
	mock := mockStore{}
	service := NewUserService(mock)

	t.Run("success", func(t *testing.T) {
		user, err := service.GetUser("vaidehi")
		assert.NilError(t, err)
		assert.DeepEqual(t, user, &models.User{Name: "vaidehi", Age: 20, Phone: 823832789, Email: "vai@c.com"})
	})

	t.Run("failure - user not found", func(t *testing.T) {
		_, err := service.GetUser("unknown")
		assert.ErrorContains(t, err, "user not found")
	})
}

func TestUserService_AddUser(t *testing.T) {
	mock := mockStore{}
	service := NewUserService(mock)

	t.Run("success", func(t *testing.T) {
		err := service.AddUser(&models.User{Name: "newUser", Age: 25, Phone: 9876543210, Email: "new@example.com"})
		assert.NilError(t, err)
	})

	t.Run("failure - duplicate user", func(t *testing.T) {
		err := service.AddUser(&models.User{Name: "duplicateUser", Age: 25, Phone: 9876543210, Email: "duplicate@example.com"})
		assert.ErrorContains(t, err, "user already exists")
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	mock := mockStore{}
	service := NewUserService(mock)

	t.Run("success", func(t *testing.T) {
		err := service.UpdateUser("existingUser", &models.User{Name: "existingUser", Age: 35, Phone: 1112223333, Email: "update@example.com"})
		assert.NilError(t, err)
	})

	t.Run("failure user not found", func(t *testing.T) {
		err := service.UpdateUser("nonExistingUser", &models.User{Name: "nonExistingUser", Age: 35, Phone: 1112223333, Email: "update@example.com"})
		assert.ErrorContains(t, err, "user not found")
	})
}

func TestUserService_DeleteUser(t *testing.T) {
	mock := mockStore{}
	service := NewUserService(mock)

	t.Run("success", func(t *testing.T) {
		err := service.DeleteUser("existingUser")
		assert.NilError(t, err)
	})

	t.Run("failure - user not found", func(t *testing.T) {
		err := service.DeleteUser("nonExistingUser")
		assert.ErrorContains(t, err, "user not found")
	})
}
