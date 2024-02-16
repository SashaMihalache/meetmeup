package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/sashamihalache/meetmeup/models"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUserById(id string) (*models.User, error) {
	return u.GetUserByField("id", id)
}

func (u *UsersRepo) GetUserByEmail(email string) (*models.User, error) {
	return u.GetUserByField("email", email)
}

func (u *UsersRepo) GetUserByUsername(username string) (*models.User, error) {
	return u.GetUserByField("username", username)
}

// helper func
func (u *UsersRepo) GetUserByField(field, value string) (*models.User, error) {
	var user models.User

	err := u.DB.Model(&user).Where(field+" = ?", value).First()

	return &user, err
}

func (u *UsersRepo) CreateUser(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}
