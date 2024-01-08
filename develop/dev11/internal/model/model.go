package model

import "dev11/utils"

type Event struct {
	EventID     int              `json:"event_id" db:"event_id"`
	UserID      int              `json:"user_id" db:"user_id"`
	Description string           `json:"description" db:"description"`
	Date        utils.CustomDate `json:"date" db:"date"`
}
