package domain

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Type      string
	TimeStamp time.Time
}
