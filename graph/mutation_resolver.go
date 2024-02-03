package graph

import (
	"context"
	"errors"

	"github.com/sashamihalache/meetmeup/graph/model"
	"github.com/sashamihalache/meetmeup/models"
)

type mutationResolver struct{ *Resolver }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := m.MeetupsRepo.GetById(id)

	if err != nil || meetup == nil {
		return nil, errors.New("meetup not found")
	}

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is too short")
		}
		meetup.Name = *input.Name
	}

	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("description is too short")
		}
		meetup.Description = *input.Description
	}

	return m.MeetupsRepo.UpdateMeetup(meetup)
}

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
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
