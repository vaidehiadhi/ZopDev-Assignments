package service

import "github.com/vaidehiadhi/threeLayerArc/models"

type StoreInterface interface {
	GetUser(name string) (*models.User, error)
	AddUser(user *models.User) error
	UpdateUser(name string, user *models.User) error
	DeleteUser(name string) error
}
