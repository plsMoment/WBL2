package service

import (
	"dev11/internal/database"
	"dev11/internal/model"
	"dev11/utils"
)

type EventManipulator interface {
	CreateEvent(event model.Event) (int, error)
	UpdateEvent(event model.Event) error
	DeleteEvent(event model.Event) error
	EventsForDay(userID int, day utils.CustomDate) ([]model.Event, error)
	EventsForWeek(userID int, day utils.CustomDate) ([]model.Event, error)
	EventsForMonth(userID int, day utils.CustomDate) ([]model.Event, error)
}

type Service struct {
	s database.IStorage
}

func NewService(s database.IStorage) *Service {
	return &Service{s}
}

func (srv *Service) CreateEvent(event model.Event) (int, error) {
	return srv.s.CreateEvent(event)
}

func (srv *Service) UpdateEvent(event model.Event) error {
	return srv.s.UpdateEvent(event)
}

func (srv *Service) DeleteEvent(event model.Event) error {
	return srv.s.DeleteEvent(event.EventID)
}

func (srv *Service) EventsForDay(userID int, day utils.CustomDate) ([]model.Event, error) {
	return srv.s.GetForDay(userID, day)
}

func (srv *Service) EventsForWeek(userID int, day utils.CustomDate) ([]model.Event, error) {
	end := utils.GetWeekRange(day)
	return srv.s.GetForRange(userID, day, end)
}

func (srv *Service) EventsForMonth(userID int, day utils.CustomDate) ([]model.Event, error) {
	end := utils.GetMonthRange(day)
	return srv.s.GetForRange(userID, day, end)
}
