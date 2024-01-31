package routes

import (
	"Service_Photo/iternal/api"

	"github.com/gin-gonic/gin"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	photo := api.NewPhotos()
	authRequired := r.Group("/")

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRequired.StaticFS("/static", http.Dir("static"))
	authRequired.POST("/", photo.Post)
	authRequired.DELETE("/:id", photo.Delete)

	return r
}
