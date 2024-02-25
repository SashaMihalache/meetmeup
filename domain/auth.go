package domain

import (
	"context"
	"errors"
	"log"

	"github.com/sashamihalache/meetmeup/models"
)

func (d *Domain) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := d.UsersRepo.GetUserByEmail(input.Email)

	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)

	if err != nil {
		return nil, ErrBadCredentials
	}

	token, err := user.GenToken()

	if err != nil {
		return nil, ErrSomethingWentWrong
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (d *Domain) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := d.UsersRepo.GetUserByEmail(input.Email)

	if err == nil {
		return nil, errors.New("email already in use")
	}

	_, err = d.UsersRepo.GetUserByUsername(input.Username)

	if err == nil {
		return nil, errors.New("username already in use")
	}

	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)

	if err != nil {
		log.Printf("error hashing password: %v", err)
		return nil, ErrSomethingWentWrong
	}

	// create a verification code with a transaction

	tx, err := d.UsersRepo.DB.Begin()

	if err != nil {
		log.Printf("error starting transaction: %v", err)
		return nil, ErrSomethingWentWrong
	}

	defer tx.Rollback()

	if _, errr := d.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("error creating user: %v", errr)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error committing transaction: %v", err)
		return nil, err
	}

	token, err := user.GenToken()

	if err != nil {
		log.Printf("error generating token: %v", err)
		return nil, ErrSomethingWentWrong
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
