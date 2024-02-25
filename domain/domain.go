package domain

import (
	"errors"

	"github.com/sashamihalache/meetmeup/models"
	"github.com/sashamihalache/meetmeup/postgres"
)

var (
	ErrBadCredentials     = errors.New("invalid credentials")
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrUnauthenticated    = errors.New("not authenticated")
	ErrForbidden          = errors.New("unauthorized")
)

type Domain struct {
	UsersRepo   postgres.UsersRepo
	MeetupsRepo postgres.MeetupsRepo
}

func NewDomain(usersRepo postgres.UsersRepo, meetupsRepo postgres.MeetupsRepo) *Domain {
	return &Domain{
		UsersRepo:   usersRepo,
		MeetupsRepo: meetupsRepo,
	}
}

type Ownable interface {
	isOwner(user *models.User) bool
}

func checkOwnership(o Ownable, user *models.User) bool {
	return o.isOwner(user)
}
