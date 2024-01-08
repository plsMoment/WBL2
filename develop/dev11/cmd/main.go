package main

import (
	"context"
	"dev11/internal/config"
	"dev11/internal/database"
	"dev11/internal/service"
	"dev11/internal/transport"
	"dev11/internal/transport/middleware"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	s, err := database.NewStorage(cfg)
	defer s.Close()

	if err != nil {
		log.Fatalf("error during connect to database: %v", err)
	}

	eventService := service.NewService(s)
	h := transport.NewHandler(eventService)
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)
	mux.HandleFunc("/events_for_day", h.GetEventsForDay)
	mux.HandleFunc("/events_for_week", h.GetEventsForWeek)
	mux.HandleFunc("/events_for_month", h.GetEventsForMonth)

	handler := middleware.Logging(mux)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    cfg.Host + ":" + cfg.Port,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("Server stopped")
			} else {
				log.Fatalf("error during server shutdown: %v", err)
			}
		}
	}()

	log.Println("Server started")

	<-done
	log.Println("Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("stopping server failed: %v", err)
	}
}
