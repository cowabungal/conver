package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PhotoRepository struct {
	db *sqlx.DB
}

func NewPhotoRepository(db *sqlx.DB) *PhotoRepository {
	return &PhotoRepository{db: db}
}

func (r *PhotoRepository) GetAllUserId() (*[]int, error) {
	usersId := make([]int, 0)

	query := fmt.Sprintf("SELECT user_id FROM %s",
		usersTable)
	err := r.db.Select(&usersId, query)
	if err != nil {
		return nil, errors.New("PhotoRepository: GetAllUserId: Select user_id: " + err.Error())
	}

	return &usersId, nil
}

func (r *PhotoRepository) SetAlbumId(userId int, albumId string) error {
	return nil
}

func (r *PhotoRepository) GetAlbumId(userId int) string {
	return ""
}

func (r *PhotoRepository) DeleteAlbumId(albumId string) error {
	return nil
}

