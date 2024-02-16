package graph

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/sashamihalache/meetmeup/models"
)

type mutationResolver struct{ *Resolver }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (m *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := m.UsersRepo.GetUserByEmail(input.Email)

	if err == nil {
		return nil, errors.New("email already in use")
	}

	_, err = m.UsersRepo.GetUserByUsername(input.Username)

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
		return nil, errors.New("something went wrong")
	}

	// create a verification code with a transaction

	tx, err := m.UsersRepo.DB.Begin()

	if err != nil {
		log.Printf("error starting transaction: %v", err)
		return nil, errors.New("something went wrong")
	}

	defer tx.Rollback()

	if _, errr := m.UsersRepo.CreateUser(tx, user); err != nil {
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
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name is too short")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description is too short")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}

	return r.MeetupsRepo.CreateMeetup(meetup)
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := m.MeetupsRepo.GetById(id)

	if err != nil || meetup == nil {
		return nil, errors.New("meetup not found")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is too short")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}

	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("description is too short")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no new data")
	}

	meetup, err = m.MeetupsRepo.UpdateMeetup(meetup)

	if err != nil {
		return nil, fmt.Errorf("error updating meetup: %v", err)
	}

	return meetup, nil
}

func (m *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := m.MeetupsRepo.GetById(id)

	if err != nil || meetup == nil {
		return false, errors.New("meetup not found")
	}

	_, err = m.MeetupsRepo.DeleteMeetup(id)

	if err != nil {
		return false, fmt.Errorf("error deleting meetup: %v", err)
	}

	return true, nil
}
