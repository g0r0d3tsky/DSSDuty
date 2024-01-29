package domain

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	Type      string    `json:"type"`
	TimeStamp time.Time `json:"time_stamp"`
}
