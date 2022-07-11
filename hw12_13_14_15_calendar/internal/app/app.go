package app

import (
	"context"
	internalconfig "github.com/cybertmt/OTUS-GO/hw12_13_14_15_calendar/internal/config"
	memorystorage "github.com/cybertmt/OTUS-GO/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/cybertmt/OTUS-GO/hw12_13_14_15_calendar/internal/storage/sql"
	"log"
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

func CreateStorage(ctx context.Context, config internalconfig.Config) (Storage, error) {
	var storage Storage
	var err error
	switch config.Storage.Type {
	case internalconfig.InMemory:
		storage = memorystorage.New()
	case internalconfig.SQL:
		storage, err = sqlstorage.New(ctx, config.Storage.Dsn).Connect(ctx)
		if err != nil {
			return nil, err
		}
	default:
		log.Fatalf("Unknown storage type: %s\n", config.Storage.Type)
	}
	return storage, nil
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
