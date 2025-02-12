package models

import "time"

// User model represents a user in the 'users' table
type User struct {
    ID        uint      `json:"id"`
    Username  string    `json:"username"`
    Password  string    `json:"-` // Password field will be excluded from response
    CreatedAt time.Time `json:"created_at"`
}
