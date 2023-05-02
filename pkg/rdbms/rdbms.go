package rdbms

import (
	"database/sql"
	"errors"
	"strings"
)

type RDBMS interface {
	MigrateUp(source string) error

	MigrateDown(source string) error

	Create(query string, args []interface{}) (uint64, error)

	Read(query string, args []interface{}, dest ...interface{}) error

	Update(query string, args []interface{}) error

	Delete(query string, args []interface{}) error
}

type rdbms struct {
	db *sql.DB
}

var (
	ErrPrepareStatement = "error when tying to prepare statement"

	ErrCreate    = "error when tying to create entry"
	ErrDuplicate = "entry exists"

	ErrRead         = "error when tying to read entry"
	ErrReadNotFound = "there is no entry with provided arguments"

	ErrUpdate = "error when tying to update entry"

	ErrDelete = "error when tying to delete entry"
)

func (db *rdbms) Create(query string, args []interface{}) (uint64, error) {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return 0, errors.New(ErrPrepareStatement)
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(args)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return 0, errors.New(ErrDuplicate)
		}

		return 0, errors.New(ErrCreate)
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return 0, errors.New(ErrCreate)
	}

	return uint64(id), nil
}

func (db *rdbms) Read(query string, args []interface{}, dest ...interface{}) error {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return errors.New(ErrPrepareStatement)
	}
	defer stmt.Close()

	result := stmt.QueryRow(args)
	err = result.Scan(dest)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(ErrReadNotFound)
		}

		return errors.New(ErrRead)
	}

	return nil
}

func (db *rdbms) Update(query string, args []interface{}) error {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return errors.New(ErrPrepareStatement)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(args); err != nil {
		return errors.New(ErrUpdate)
	}

	return nil
}

func (db *rdbms) Delete(query string, args []interface{}) error {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return errors.New(ErrPrepareStatement)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(args); err != nil {
		return errors.New(ErrDelete)
	}

	return nil
}
