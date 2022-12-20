package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"puppy/Utility"
	"puppy/createphoto/common/httpResponse"
	"puppy/createphoto/photocr"
	"puppy/service"
)

type PhotoController interface {
	GetPhotoList(c echo.Context) error
	GetPhoto(c echo.Context) error
	CreatePhoto(c echo.Context) error
	EditPhoto(c echo.Context) error
	DeletePhoto(c echo.Context) error
	LikePhoto(c echo.Context) error
	UploadPhoto(c echo.Context) error
}

type photoController struct {
}

func NewPhotoController() PhotoController {
	return photoController{}
}

func (nc photoController) GetPhotoList(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

	photoservice := service.NewPhotoService()
	newsList, err := photoservice.GetPhotoList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(newsList))
}

func (nc photoController) GetPhoto(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetNewsId := apiContext.Param("id")

	photoService := service.NewPhotoService()
	news, err := photoService.GetPhotoId(targetNewsId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.NotFoundResponse(nil, "photocr not found"))
	}

	photoService.AddVisitCount(targetNewsId)

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(news))
}

func (nc photoController) CreatePhoto(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	newphoto := new(photocr.CreatePhoto)

	if err := apiContext.Bind(newphoto); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data not found"))
	}

	if err := c.Validate(newphoto); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	file, err := apiContext.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("image not found"))
	}

	photoService := service.NewPhotoService()
	newphotoId, err := photoService.CreateNewPhoto(*newphoto, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: newphotoId,
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}

func (nc photoController) EditPhoto(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetphotoId := apiContext.Param("id")

	editphoto := new(photocr.EditPhoto)

	if err := apiContext.Bind(editphoto); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data not found"))
	}

	if err := c.Validate(editphoto); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	file, err := apiContext.FormFile("file")

	editphoto.Id = targetphotoId

	photoservice := service.NewPhotoService()

	if !photoservice.IsPhotoExist(targetphotoId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err = photoservice.EditPhoto(*editphoto, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(nil))
}

func (nc photoController) DeletePhoto(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetNewsId := apiContext.Param("id")

	photoservice := service.NewPhotoService()
	if !photoservice.IsPhotoExist(targetNewsId) {
		return c.JSON(http.StatusNotFound, httpResponse.SuccessResponse("Data not found"))
	}

	err := photoservice.DeletePhoto(targetNewsId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(nil))
}

func (nc photoController) LikePhoto(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetNewsId := apiContext.Param("id")

	photoservice := service.NewPhotoService()

	if !photoservice.IsPhotoExist(targetNewsId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := photoservice.AddLike(targetNewsId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(nil))
}

func (nc photoController) UploadPhoto(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	file, err := apiContext.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	des, err := os.Create(file.Filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer des.Close()

	_, err = io.Copy(des, src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}
