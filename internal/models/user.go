package models

import "time"

type User struct {
	ID        string       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // Exclude password from JSON responses for security
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserDTO struct {
	Name           *string `form:"name" json:"name,omitempty"` // Omitting empty JSON fields
	ProfilePicture *string `json:"profile_picture,omitempty"`
}
