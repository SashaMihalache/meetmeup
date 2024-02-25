package graph

import (
	"context"
	"errors"

	"github.com/sashamihalache/meetmeup/models"
)

var (
	ErrInput = errors.New("Input errors")
)

type mutationResolver struct{ *Resolver }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (m *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)

	if !isValid {
		return nil, ErrInput
	}

	return m.Domain.Login(ctx, input)
}

func (m *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)

	if !isValid {
		return nil, ErrInput
	}

	return m.Domain.Register(ctx, input)
}

// CreateMeetup is the resolver for the createMeetup field.
func (m *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	return m.Domain.CreateMeetup(ctx, input)
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	return m.Domain.UpdateMeetup(ctx, id, input)
}

func (m *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	return m.Domain.DeleteMeetup(ctx, id)
}
