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
	QueryCreateUser = `INSERT INTO users(first_name, last_name, email, password) VALUES(?, ?, ?, ?);`
)

func (r *repository) CreateUser(ctx context.Context, user *models.User) error {
	if len(user.Email) == 0 || len(user.Password) == 0 {
		return errors.New("Insufficient information for user")
	}

	args := []interface{}{user.FirstName, user.LastName, user.Email, user.Password}
	id, err := r.rdbms.Create(QueryCreateUser, args)
	if err != nil {
		r.logger.Error("Error creating user", zap.Error(err))
		return err
	}

	user.Id = id
	return nil
}

const QueryFindUserById = "SELECT first_name, last_name, email, password, created_at FROM users WHERE id=?"

func (r *repository) FindUserById(ctx context.Context, id uint64) (*models.User, error) {
	user := &models.User{Id: id}

	args := []interface{}{id}
	dest := []interface{}{&user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt}
	if err := r.rdbms.Read(QueryFindUserById, args, dest); err != nil {
		r.logger.Error("Error find user by id", zap.Error(err))
		return nil, err
	}

	return user, nil
}

const QueryFindUserByEmailAndPassword = "SELECT id, first_name, last_name, created_at FROM users WHERE email=? AND password=?"

func (r *repository) FindUserByEmailAndPassword(ctx context.Context, email, password string) (*models.User, error) {
	user := &models.User{Email: email, Password: password}

	args := []interface{}{email, password}
	dest := []interface{}{&user.Id, &user.FirstName, &user.LastName, &user.CreatedAt}
	if err := r.rdbms.Read(QueryFindUserByEmailAndPassword, args, dest); err != nil {
		r.logger.Error("Error find user by email and password", zap.Error(err))
		return nil, err
	}

	return user, nil
}

const QueryFindUserByEmail = "SELECT id, first_name, last_name, created_at FROM users WHERE email=?"

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{Email: email}

	args := []interface{}{email}
	dest := []interface{}{&user.Id, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt}
	if err := r.rdbms.Read(QueryFindUserByEmail, args, dest); err != nil {
		r.logger.Error("Error find user by email", zap.Error(err))
		return nil, err
	}

	return user, nil
}

const QueryUpdateUser = "UPDATE users SET first_name=?, last_name=?, password=? WHERE id=?;"

// UpdateUser will only updates the first_name and last_name or password
func (r *repository) UpdateUser(ctx context.Context, user *models.User) error {
	args := []interface{}{user.FirstName, user.LastName, user.Password, user.Id}
	if err := r.rdbms.Update(QueryUpdateUser, args); err != nil {
		r.logger.Error("Error updating user", zap.Any("user", user), zap.Error(err))
		return err
	}

	return nil
}

const QueryDeleteUser = "DELETE FROM users WHERE id=?;"

func (r *repository) DeleteUser(ctx context.Context, user *models.User) error {
	args := []interface{}{user.Id}
	if err := r.rdbms.Delete(QueryDeleteUser, args); err != nil {
		r.logger.Error("Error deleting user", zap.Any("user", user), zap.Error(err))
		return err
	}

	return nil
}
