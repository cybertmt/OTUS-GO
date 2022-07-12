package production

import (
	"context"
	"github.com/cybertmt/OTUS-GO/hw12_13_14_15_calendar/internal/app"
	internalconfig "github.com/cybertmt/OTUS-GO/hw12_13_14_15_calendar/internal/config"
	memorystorage "github.com/cybertmt/OTUS-GO/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/cybertmt/OTUS-GO/hw12_13_14_15_calendar/internal/storage/sql"
	"log"
)

func CreateStorage(ctx context.Context, config internalconfig.Config) (app.Storage, error) {
	var storage app.Storage
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
