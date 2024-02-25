package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/sashamihalache/meetmeup/middleware"
	"github.com/sashamihalache/meetmeup/models"
)

// CreateMeetup is the resolver for the createMeetup field.
func (d *Domain) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, ErrUnauthenticated
	}

	if len(input.Name) < 3 {
		return nil, errors.New("name is too short")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description is too short")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}

	return d.MeetupsRepo.CreateMeetup(meetup)
}

func (d *Domain) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, ErrUnauthenticated
	}

	meetup, err := d.MeetupsRepo.GetById(id)

	if err != nil || meetup == nil {
		return nil, errors.New("meetup not found")
	}

	if !meetup.IsOwner(currentUser) {
		return nil, ErrForbidden
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

	meetup, err = d.MeetupsRepo.UpdateMeetup(meetup)

	if err != nil {
		return nil, fmt.Errorf("error updating meetup: %v", err)
	}

	return meetup, nil
}

func (d *Domain) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return false, ErrUnauthenticated
	}

	meetup, err := d.MeetupsRepo.GetById(id)

	if err != nil || meetup == nil {
		return false, errors.New("meetup not found")
	}

	if !meetup.IsOwner(currentUser) {
		return false, ErrForbidden
	}

	_, err = d.MeetupsRepo.DeleteMeetup(id)

	if err != nil {
		return false, fmt.Errorf("error deleting meetup: %v", err)
	}

	return true, nil
}
