package test_ci

import (
	"context"
	postgresorm "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Type     string
	URL      string
	Database string
}

// New opengauss datastore instance
func New(ctx context.Context, cfg Config) error {

	_, err := gorm.Open(postgresorm.Open(cfg.URL), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return nil
}
