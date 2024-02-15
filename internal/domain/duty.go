package domain

import (
	"github.com/google/uuid"
	"time"
)

type Duty struct {
	Id      uuid.UUID `json:"id"`
	Date    time.Time `json:"date"`
	UserId  UserId    `json:"user_id"`
	Amount  int       `json:"amount"`
	Picking Picking
}

// todo fix
type UserId struct {
	First  uuid.UUID `json:"first"`
	Second uuid.UUID `json:"second"`
}

type Picking struct {
	PickingTimeStart time.Time `json:"picking_time_start"`
	PickingTimeEnd   time.Time `json:"picking_time_end"`
}
