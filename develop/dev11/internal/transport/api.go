package transport

import (
	"dev11/internal/model"
	"dev11/internal/service"
	"dev11/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func errorResponse(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(ErrorResponse{err.Error()})
	w.Write(b)
}

type Handler struct {
	srv service.EventManipulator
}

func NewHandler(srv service.EventManipulator) *Handler {
	return &Handler{srv}
}

type IntResponse struct {
	EventID int `json:"event_id"`
}

func (h *Handler) CreateEvent(w http.ResponseWriter, req *http.Request) {
	var event model.Event
	err := json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	eventID, err := h.srv.CreateEvent(event)
	if err != nil {
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(IntResponse{eventID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

type StringResponse struct {
	Result string `json:"result"`
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, req *http.Request) {
	var event model.Event
	err := json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err = h.srv.UpdateEvent(event); err != nil {
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(StringResponse{"Success"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *Handler) DeleteEvent(w http.ResponseWriter, req *http.Request) {
	var event model.Event
	err := json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err = h.srv.UpdateEvent(event); err != nil {
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(StringResponse{"Success"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

type SliceResponse struct {
	Result []model.Event `json:"result"`
}

func (h *Handler) GetEventsForDay(w http.ResponseWriter, req *http.Request) {
	userIDstr := req.URL.Query().Get("user_id")
	dateStr := req.URL.Query().Get("date")

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	date, err := utils.GetDate(dateStr)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	events, err := h.srv.EventsForDay(userID, date)
	if err != nil {
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(SliceResponse{events})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetEventsForWeek(w http.ResponseWriter, req *http.Request) {
	userIDstr := req.URL.Query().Get("user_id")
	dateStr := req.URL.Query().Get("date")

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	date, err := utils.GetDate(dateStr)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	events, err := h.srv.EventsForWeek(userID, date)
	if err != nil {
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(SliceResponse{events})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetEventsForMonth(w http.ResponseWriter, req *http.Request) {
	userIDstr := req.URL.Query().Get("user_id")
	dateStr := req.URL.Query().Get("date")

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	date, err := utils.GetDate(dateStr)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	events, err := h.srv.EventsForMonth(userID, date)
	if err != nil {
		errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(SliceResponse{events})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
