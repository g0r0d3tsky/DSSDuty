package domain

import (
	"github.com/google/uuid"
	"time"
)

type Duty struct {
	Id     uuid.UUID
	Date   time.Time
	UserId uuid.UUID
}
