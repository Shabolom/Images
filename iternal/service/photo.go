package service

import (
	"Service_Photo/iternal/domain"
	"Service_Photo/iternal/repositori"
	"github.com/gofrs/uuid"
	"mime/multipart"
)

type PhotoServise struct{}

func NewPhotoServise() *PhotoServise {
	return &PhotoServise{}
}

var photoRepo = repositori.NewPhotoRepo()

func (ps *PhotoServise) Post(image *multipart.FileHeader, path string) error {
	id, _ := uuid.NewV4()

	photoEntite := domain.Photos{
		Name:      image.Filename,
		Path:      path,
		Size:      image.Size,
		PathToMin: path,
	}
	photoEntite.ID = id

	err := photoRepo.Post(photoEntite)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PhotoServise) Del(key string) error {
	err := photoRepo.Del(key)
	if err != nil {
		return err
	}
	return nil
}
