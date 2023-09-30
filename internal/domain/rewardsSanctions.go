package domain

import (
	"github.com/google/uuid"
	"time"
)

type RewardsSanctions struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Rewards   int
	Sanctions int
	Info      string
	DateTime  time.Time
}
