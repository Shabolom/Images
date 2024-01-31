package api

import (
	"Service_Photo/iternal/service"
	"Service_Photo/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Photo struct{}

func NewPhotos() *Photo {
	return &Photo{}
}

var photoService = service.NewPhotoServise()

// Post заносим в базу данных данные jpg и создаем локалбно директорию с датой когда создана и создаем там файл куда копируем изначальный файл
// @Summary  заносим в базу данные файла и сохраняем локально сам файл, а так же создаем каждый день новую директорию
// @Accept   jpeg
// @Produce  json
// @Tags     photo
// @Param image formData file true "Body with file"
// @Success  201  {string}  string    "ok"
// @Failure  400  {object}  models.Error
// @Failure  404  {object}  models.Error
// @Router / [post]
func (p *Photo) Post(c *gin.Context) {
	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		log.WithField("component", "rest").Warn(err)
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	defer file.Close()

	t := time.Now()
	timeString := t.Format("2006-01-02")

	err = os.MkdirAll("static/"+timeString, 0666)
	if err != nil {
		log.WithField("component", "rest").Warn(err)
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	dir := "static/" + timeString + "/"

	fileName := fmt.Sprintf("%d.jpg", time.Now().Unix())

	filePath := filepath.Join(dir, fileName)

	targetFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.WithField("component", "rest").Warn(err)
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	defer targetFile.Close()

	if err = photoService.Post(handler, filePath); err != nil {
		log.WithField("component", "rest").Warn(err)
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	_, err = io.Copy(targetFile, file)
	if err != nil {
		log.WithField("component", "rest").Warn(err)
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, "ваш файл сохранен")
}

// Delete делаем soft мягкое удаление айди из базы данных
// @Summary  делаем soft мягкое удаление айди из базы данных
// @Accept   json
// @Tags     photo
// @Param id path string true "id of photo"
// @Success  201  {string}  string    "ok"
// @Failure  400  {object}  models.Error
// @Failure  404  {object}  models.Error
// @Router   /{id} [delete]
func (p *Photo) Delete(c *gin.Context) {
	id := c.Param("id")

	err := photoService.GetID(id)

	if err != nil {
		log.WithField("component", "rest").Warn(err)
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, "ваш файл успешно удален")
}
