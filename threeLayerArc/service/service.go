package service

import (
	"github.com/vaidehiadhi/threeLayerArc/models"
)

type userService struct {
	store StoreInterface
}

func NewUserService(store StoreInterface) userService {
	return userService{store: store}
}

func (s userService) GetUser(name string) (*models.User, error) {
	return s.store.GetUser(name)
}

func (s userService) AddUser(user *models.User) error {
	return s.store.AddUser(user)
}

func (s userService) UpdateUser(name string, user *models.User) error {
	return s.store.UpdateUser(name, user)
}

func (s userService) DeleteUser(name string) error {
	return s.store.DeleteUser(name)
}
