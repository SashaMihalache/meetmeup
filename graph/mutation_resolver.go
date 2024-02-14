package graph

import (
	"context"
	"errors"
	"fmt"

	"github.com/sashamihalache/meetmeup/models"
)

type mutationResolver struct{ *Resolver }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

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
