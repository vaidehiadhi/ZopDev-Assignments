package service

import (
	"github.com/vaidehiadhi/threeLayerArc/models"
	"github.com/vaidehiadhi/threeLayerArc/store"
)

type UserService struct {
	store *store.Store
}

func NewUserService(store *store.Store) *UserService {
	return &UserService{store: store}
}

func (s *UserService) GetUser(name string) (*models.User, error) {
	return s.store.GetUser(name)
}

func (s *UserService) AddUser(user *models.User) error {
	return s.store.AddUser(user)
}

func (s *UserService) UpdateUser(name string, user *models.User) error {
	return s.store.UpdateUser(name, user)
}

func (s *UserService) DeleteUser(name string) error {
	return s.store.DeleteUser(name)
}
