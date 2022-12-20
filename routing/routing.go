package routing

import (
	"github.com/labstack/echo/v4"
	"puppy/controller"
)

func SetRouting(e *echo.Echo) error {
	RoutePhotoController(e)

	return nil
}

func RoutePhotoController(e *echo.Echo) {
	photoController := controller.NewPhotoController()

	photogroup := e.Group("photo")

	photogroup.GET("/getList", photoController.GetPhotoList)
	photogroup.GET("/:id", photoController.GetPhoto)
	photogroup.GET("/:id/Like", photoController.LikePhoto)
	photogroup.POST("/Create", photoController.CreatePhoto)
	photogroup.POST("/Edit/:id", photoController.EditPhoto)
	photogroup.DELETE("/Delete/:id", photoController.DeletePhoto)
	photogroup.POST("/uploadphoto", photoController.UploadPhoto)
}
