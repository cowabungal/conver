package service

import (
	"conver/pkg/repository"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type PhotoService struct {
	repo *repository.Repository
}

func NewPhotoService(repo *repository.Repository) *PhotoService {
	return &PhotoService{repo: repo}
}

func (s *PhotoService) Save(userId int, reader io.Reader) (int, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return 0, err
	}

	readDirectory, _ := os.Open(fmt.Sprintf("./assets/%d/", userId))
	allFiles, _ := readDirectory.Readdir(0)
	n := len(allFiles) + 1

	return n, ioutil.WriteFile(fmt.Sprintf("./assets/%d/%d.jpg", userId, n), body, 0o666)
}

func CreateDir(userId int) error {
	return os.Mkdir(fmt.Sprintf("./assets/%d", userId), 0o777)
}

func (s *PhotoService) GetPdf(userId int) error {
	GoPdf(fmt.Sprintf("./assets/%d/", userId))
	return nil
}


func (s *PhotoService) SetAlbumId(userId int, albumId string) error {
	return s.repo.SetAlbumId(userId, albumId)
}

func (s *PhotoService) GetAlbumId(userId int) string {
	return s.repo.GetAlbumId(userId)
}

func (s *PhotoService) DeleteAlbumId(albumId string) error {
	return s.repo.DeleteAlbumId(albumId)
}
