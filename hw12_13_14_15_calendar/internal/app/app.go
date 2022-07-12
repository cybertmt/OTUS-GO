package app

import (
	"context"
	"net/http"
	"time"

	"github.com/cybertmt/OTUS-GO/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

type App struct {
	Logger  Logger
	Storage Storage
}

type Logger interface {
	Debug(format string, params ...interface{})
	Info(format string, params ...interface{})
	Warn(format string, params ...interface{})
	Error(format string, params ...interface{})
	LogHTTPRequest(r *http.Request, code, length int)
}

type Storage interface {
	Create(e storage.Event) error
	Update(e storage.Event) error
	Delete(id uuid.UUID) error
	FindAll() ([]storage.Event, error)
	FindAtMonth(dayStart time.Time) ([]storage.Event, error)
}

func New(logger Logger, storage Storage) *App {
	return &App{
		Logger:  logger,
		Storage: storage,
	}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
