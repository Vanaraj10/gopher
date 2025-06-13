package models

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"` // Password should not be exposed in API responses
	CreatedAt string `json:"created_at"`
	Bio 	 string `json:"bio,omitempty"` // Optional field for user bio
}