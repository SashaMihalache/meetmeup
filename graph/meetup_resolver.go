package graph

import (
	"context"

	"github.com/sashamihalache/meetmeup/models"
)

type meetupResolver struct{ *Resolver }

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return r.UsersRepo.GetUserById(obj.UserID)
}

// Meetup returns MeetupResolver implementation.
func (r *Resolver) Meetup() MeetupResolver { return &meetupResolver{r} }
