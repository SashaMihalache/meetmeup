// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MeetupFilter struct {
	Name *string `json:"name,omitempty"`
}

type Mutation struct {
}

type NewMeetup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Query struct {
}

type RegisterInput struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
}

type UpdateMeetup struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
