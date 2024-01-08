package database

import (
	"dev11/internal/model"
	"dev11/utils"
)

var (
	createStmt     = "INSERT INTO events (user_id, description, date) VALUES (?, ?, ?)"
	updateStmt     = "UPDATE events SET user_id=?, description=?, date=? WHERE event_id=?"
	deleteStmt     = "DELETE FROM events WHERE event_id=?"
	getStmt        = "SELECT * FROM events WHERE user_id=? and date=?"
	getInRangeStmt = "SELECT * FROM events WHERE user_id=? and date BETWEEN ? AND ?"
)

func (s *Storage) CreateEvent(event model.Event) (int, error) {
	res, err := s.db.Exec(createStmt, event.UserID, event.Description, event.Date)
	if err != nil {
		return 0, err
	}

	eventID, err := res.LastInsertId()

	return int(eventID), err
}

func (s *Storage) UpdateEvent(event model.Event) error {
	_, err := s.db.Exec(updateStmt, event.UserID, event.Description, event.Date, event.EventID)
	return err
}

func (s *Storage) DeleteEvent(eventID int) error {
	_, err := s.db.Exec(deleteStmt, eventID)
	return err
}

func (s *Storage) GetForDay(userID int, day utils.CustomDate) ([]model.Event, error) {
	var events []model.Event
	err := s.db.Select(&events, getStmt, userID, day)
	return events, err
}

func (s *Storage) GetForRange(userID int, begin, end utils.CustomDate) ([]model.Event, error) {
	var events []model.Event
	err := s.db.Select(&events, getInRangeStmt, userID, begin, end)
	return events, err
}
