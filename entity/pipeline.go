package entity

import "github.com/google/uuid"

type Pipeline struct {
	ID   uuid.UUID
	Name string
	// Status can be ENABLED | DISABLED
	Status string
}
