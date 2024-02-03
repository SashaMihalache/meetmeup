package graph

import (
	"context"

	"github.com/sashamihalache/meetmeup/models"
)

type meetupResolver struct{ *Resolver }

// Meetup returns MeetupResolver implementation.
func (r *Resolver) Meetup() MeetupResolver { return &meetupResolver{r} }

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return getUserLoader(ctx).Load(obj.UserID)
}
