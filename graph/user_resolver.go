package graph

import (
	"context"

	"github.com/sashamihalache/meetmeup/models"
)

type userResolver struct{ *Resolver }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// Meetups is the resolver for the meetups field.
func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetupsForUser(obj)
}
