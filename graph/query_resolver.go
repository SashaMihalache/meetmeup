package graph

import (
	"context"

	"github.com/sashamihalache/meetmeup/models"
)

type queryResolver struct{ *Resolver }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context, filter *models.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetups(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UsersRepo.GetUserById(id)
}
