// go:generate go run github.com/99designs/gqlgen -v

package graph

import "github.com/sashamihalache/meetmeup/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UsersRepo   postgres.UsersRepo
}
