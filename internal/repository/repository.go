package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Repository interface {
	MigrateUp(context.Context) error
	MigrateDown(context.Context) error
}

type repository struct {
	logger *zap.Logger
	db     *sql.DB
}

func New(cfg *Config, lg *zap.Logger) Repository {
	r := &repository{logger: lg}

	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		lg.Panic("Error openning postgresql connection", zap.Error(err))
	}
	r.db = db

	return r
}

func (r *repository) MigrateUp(ctx context.Context) error {
	do := func(m *migrate.Migrate) error { return m.Up() }
	return r.migrate(ctx, do)
}

func (r *repository) MigrateDown(ctx context.Context) error {
	do := func(m *migrate.Migrate) error { return m.Down() }
	return r.migrate(ctx, do)
}

func (r *repository) migrate(ctx context.Context, do func(*migrate.Migrate) error) error {
	instance, err := postgres.WithInstance(r.db, &postgres.Config{})
	if err != nil {
		errorString := "Error creating migrate instance"
		r.logger.Error(errorString, zap.Error(err))
		return errors.New(errorString)
	}

	source := "file://internal/repository/migrations"
	migrator, err := migrate.NewWithDatabaseInstance(source, "postgres", instance)
	if err != nil {
		errorString := "Error loading migration files"
		r.logger.Error(errorString, zap.Error(err))
		return errors.New(errorString)
	}

	if err := do(migrator); err != nil && err != migrate.ErrNoChange {
		errorString := "Error doing migrations"
		r.logger.Error(errorString, zap.Error(err))
		return errors.New(errorString)
	}

	return nil
}
