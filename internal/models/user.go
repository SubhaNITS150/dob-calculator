package models

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	DOB  string    `json:"dob"`
	Age  int       `json:"age,omitempty"`
}
