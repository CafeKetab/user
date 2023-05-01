package repository

import (
	"context"
	"errors"

	"github.com/CafeKetab/user/internal/models"
	"github.com/CafeKetab/user/pkg/rdbms"

	"go.uber.org/zap"
)

type Repository interface {
	MigrateUp(context.Context) error
	MigrateDown(context.Context) error
}

type repository struct {
	logger             *zap.Logger
	rdbms              rdbms.RDBMS
	migrationDirectory string
}

func New(lg *zap.Logger, rdbms rdbms.RDBMS) Repository {
	r := &repository{logger: lg, rdbms: rdbms}
	r.migrationDirectory = "file://internal/repository/migrations"

	return r
}

func (r *repository) MigrateUp(ctx context.Context) error {
	return r.rdbms.MigrateUp(r.migrationDirectory)
}

func (r *repository) MigrateDown(ctx context.Context) error {
	return r.rdbms.MigrateDown(r.migrationDirectory)
}

const (
	QueryCreateUserInformation = `INSERT INTO user_informations(first_name, last_name) VALUES(?, ?);`
	QueryCreateUser            = `INSERT INTO users(email, password, user_information_id) VALUES(?, ?, ?);`
)

func (r *repository) CreateUser(ctx context.Context, user *models.User) error {
	if len(user.Email) == 0 || len(user.Password) == 0 {
		return errors.New("Insufficient information for user")
	}

	userInformationArgs := []interface{}{user.FirstName, user.LastName}
	userInformationId, err := r.rdbms.Create(QueryCreateUserInformation, userInformationArgs)
	if err != nil {
		r.logger.Error("Error creating user_information", zap.Error(err))
		return err
	}

	userArgs := []interface{}{user.Email, user.Password, userInformationId}
	userId, err := r.rdbms.Create(QueryCreateUser, userArgs)
	if err != nil {
		r.logger.Error("Error creating user", zap.Error(err))
		return err
	}

	user.Id = userId
	return nil
}
