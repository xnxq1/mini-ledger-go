package domain

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID       uuid.UUID `json:"id"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Archived bool      `json:"archived"`
}
