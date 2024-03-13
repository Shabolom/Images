package repositori

import (
	"Service_Photo/config"
	"Service_Photo/iternal/domain"
)

type PhotoRepo struct{}

func NewPhotoRepo() *PhotoRepo {
	return &PhotoRepo{}
}

func (pr *PhotoRepo) Post(domain domain.Photos) error {

	err := config.DB.Create(&domain).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *PhotoRepo) Del(key string) error {
	var photo domain.Photos

	err := config.DB.
		Where("id = ?", key).
		Delete(&photo).
		Error
	if err != nil {
		return err
	}
	return nil
}
