package models

import "time"

type User struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Nick       string    `json:"nick"`
	Email      string    `json:"email"`
	Created_At time.Time `json:"created_at"`
	Followers  []User    `json:"followers"`
	Following  []User    `json:"following"`
	Posts      []Post    `json:"posts"`
}
