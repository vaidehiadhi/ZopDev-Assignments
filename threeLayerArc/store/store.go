package store

import (
	"database/sql"
	"errors"
	"github.com/vaidehiadhi/threeLayerArc/models"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) GetUser(name string) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow("SELECT  name, age, phone, email FROM `user` WHERE name = ?", name).
		Scan(&user.Name, &user.Age, &user.Phone, &user.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &user, err
}

func (s *store) AddUser(user *models.User) error {
	_, err := s.db.Exec("INSERT INTO `user` (name, age, phone, email) VALUES (?,?,?,?)",
		user.Name, user.Age, user.Phone, user.Email)
	return err
}

func (s *store) UpdateUser(name string, user *models.User) error {
	_, err := s.db.Exec("UPDATE `user` SET age= ?, phone = ?, email = ? WHERE name = ?",
		user.Age, user.Phone, user.Email, user.Name)
	return err
}

func (s *store) DeleteUser(name string) error {
	result, err := s.db.Exec("DELETE FROM `user` WHERE name = ?", name)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
