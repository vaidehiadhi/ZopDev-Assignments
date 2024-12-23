package store

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vaidehiadhi/threeLayerArc/models"
	"gotest.tools/v3/assert"
	"regexp"
	"testing"
)

func TestStore_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name, age, phone, email FROM `user` WHERE name = ?").
		WithArgs("vai vai").
		WillReturnRows(sqlmock.NewRows([]string{"name", "age", "phone", "email"}).
			AddRow("vai vai", 20, 823832789, "vai@c.com"))

	st := NewStore(db)
	user, err := st.GetUser("vai vai")
	assert.NilError(t, err)
	assert.Equal(t, "vai vai", user.Name)
	assert.Equal(t, 20, user.Age)
	assert.Equal(t, 823832789, user.Phone)
	assert.Equal(t, "vai@c.com", user.Email)
}

func TestStore_AddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectExec("INSERT INTO `user`").
		WithArgs("vaidehi", 20, 8273819101, "vai@g.com").
		WillReturnResult(sqlmock.NewResult(1, 1))

	st := NewStore(db)
	user := &models.User{Name: "vaidehi", Age: 20, Phone: 8273819101, Email: "vai@g.com"}
	err = st.AddUser(user)
	assert.NilError(t, err)
}

func TestStore_UpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE `user` SET age= ?, phone = ?, email = ? WHERE name = ?")).
		WithArgs(22, 8289192020, "vai_new@g.com", "vai").
		WillReturnResult(sqlmock.NewResult(1, 1))

	st := NewStore(db)
	user := &models.User{
		Name:  "vai",
		Age:   22,
		Phone: 8289192020,
		Email: "vai_new@g.com",
	}
	err = st.UpdateUser("vai", user)
	assert.NilError(t, err)
}

func TestStore_DeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectExec("DELETE FROM `user` WHERE name = ?").
		WithArgs("vai vai").
		WillReturnResult(sqlmock.NewResult(1, 1))

	st := NewStore(db)
	err = st.DeleteUser("vai vai")
	assert.NilError(t, err)
}

func TestStore_DeleteUser_NoRowsAffected(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectExec("DELETE FROM `user` WHERE name = ?").
		WithArgs("nonExistentUser").
		WillReturnResult(sqlmock.NewResult(0, 0))

	st := NewStore(db)
	err = st.DeleteUser("nonExistentUser")
	assert.ErrorContains(t, err, "no rows in result set")
}
