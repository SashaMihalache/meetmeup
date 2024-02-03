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
	var result []*models.Meetup

	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			result = append(result, meetup)
		}
	}

	return result, nil
}
