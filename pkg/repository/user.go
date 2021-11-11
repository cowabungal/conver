package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(username string, userId int) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username, user_id) values ($1, $2)", usersTable)

	row := r.db.QueryRow(query, username, userId)
	err := row.Scan(&id)
	return err
}

func (r *UserRepository) State(userId int) (string, error) {
	var state string

	query := fmt.Sprintf("SELECT state from %s WHERE user_id=$1;", usersTable)
	err := r.db.Get(&state, query, userId)

	return state, err
}

func (r *UserRepository) ChangeState(userId int, newState string) (string, error) {
	var state string

	query := fmt.Sprintf("UPDATE %s SET %s=$1 WHERE user_id=$2 RETURNING state;", usersTable, stateColumn)
	err := r.db.Get(&state, query, newState, userId)

	return state, err
}

func (r *UserRepository) AddCallbackId(userId int, callbackId string) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, callback_id) values ($1, $2) RETURNING user_id", callbacksTable)

	var tmp string

	row := r.db.QueryRow(query, userId, callbackId)
	err := row.Scan(&tmp)

	return err
}

func (r *UserRepository) AddCallbackData(callbackId, callbackData string) error {
	var tmp string

	query := fmt.Sprintf("UPDATE %s SET %s=$1 WHERE callback_id=$2 RETURNING callback_data;", callbacksTable, callbackDataColumn)
	err := r.db.Get(&tmp, query, callbackData, callbackId)

	return err
}

func (r *UserRepository) GetCallbackData(userId int) (string, error) {
	var data string

	query := fmt.Sprintf("SELECT callback_data from %s WHERE user_id=$1;", callbacksTable)
	err := r.db.Get(&data, query, userId)

	return data, err
}

func (r *UserRepository) GetCallbackId(userId int) (int, error) {
	var id int

	query := fmt.Sprintf("SELECT callback_id from %s WHERE user_id=$1;", callbacksTable)
	err := r.db.Get(&id, query, userId)

	return id, err
}

func (r *UserRepository) DeleteCallback(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", callbacksTable)

	rows, err := r.db.Query(query, userId)
	rows.Close()

	return err
}
