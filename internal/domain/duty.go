package domain

import (
	"github.com/google/uuid"
	"time"
)

type Duty struct {
	Id     uuid.UUID
	Date   time.Time
	UserId UserId
}

// todo fix
type UserId struct {
	First  uuid.UUID
	Second uuid.UUID
}
