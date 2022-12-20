package service

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"puppy/Utility"
	"puppy/createphoto/photocr"
	"puppy/model/photo"
	"puppy/repository"
	"time"
)

type PhotoService interface {
	GetPhotoList() ([]photo.Photo, error)
	UploadPhoto(c echo.Context) error
	GetPhotoId(id string) (photo.Photo, error)
	CreateNewPhoto(userInput photocr.CreatePhoto, imageFile *multipart.FileHeader) (string, error)
	IsPhotoExist(id string) bool
	EditPhoto(userInput photocr.EditPhoto, imageFile *multipart.FileHeader) error
	DeletePhoto(id string) error
	AddVisitCount(id string) error
	AddLike(id string) error
}

type photoService struct {
}

func NewPhotoService() PhotoService {
	return photoService{}
}

func (photoService) GetPhotoList() ([]photo.Photo, error) {

	photoRepository := repository.NewPhotoRepository()
	photoList, err := photoRepository.GetPhotoList()
	return photoList, err
}

func (s photoService) GetPhotoId(id string) (photo.Photo, error) {
	photoRepository := repository.NewPhotoRepository()
	photoo, err := photoRepository.GetPhotoById(id)
	return photoo, err
}

func (s photoService) CreateNewPhoto(userInput photocr.CreatePhoto, imageFile *multipart.FileHeader) (string, error) {

	photoEntity := photo.Photo{
		Title:         userInput.Title,
		ImageName:     userInput.ImageName,
		Description:   userInput.Description,
		CreateDate:    time.Now(),
		CreatorUserId: userInput.CreatorUserId,
	}
	if imageFile != nil {
		src, err := imageFile.Open()
		if err != nil {
			return "", err
		}

		fileName := uuid.New().String() + filepath.Ext(imageFile.Filename)

		wd, err := os.Getwd()
		imageServerPath := filepath.Join(wd, "wwwroot", "images", fileName)

		des, err := os.Create(imageServerPath)
		if err != nil {
			return "", err
		}
		defer des.Close()

		_, err = io.Copy(des, src)
		if err != nil {
			return "", err
		}
		photoEntity.ImageName = fileName
	}

	photoRepository := repository.NewPhotoRepository()
	photoid, err := photoRepository.InsertPhoto(photoEntity)

	return photoid, err
}

func (s photoService) IsPhotoExist(id string) bool {
	newsRepository := repository.NewPhotoRepository()
	_, err := newsRepository.GetPhotoById(id)

	if err != nil {
		return false
	}

	return true
}

func (s photoService) EditPhoto(userInput photocr.EditPhoto, imageFile *multipart.FileHeader) error {

	photorepository := repository.NewPhotoRepository()
	photoentity := photo.Photo{
		Id:            userInput.Id,
		Title:         userInput.Title,
		ImageName:     userInput.ImageName,
		Description:   userInput.Description,
		CreateDate:    time.Now(),
		CreatorUserId: userInput.CreatorUserId,
	}
	if imageFile != nil {
		src, err := imageFile.Open()
		if err != nil {
			return err
		}
		oldPhoto, err := photorepository.GetPhotoById(userInput.Id)
		if err != nil {
			return err
		}

		wd, err := os.Getwd()

		if oldPhoto.ImageName != "" {
			oldImageServerPath := filepath.Join(wd, "wwwroot", "images", "photocr", oldPhoto.ImageName)
			os.Remove(oldImageServerPath)
		}

		fileName := uuid.New().String() + filepath.Ext(imageFile.Filename)

		imageServerPath := filepath.Join(wd, "wwwroot", "images", "photocr", fileName)

		des, err := os.Create(imageServerPath)
		if err != nil {
			return err
		}
		defer des.Close()

		_, err = io.Copy(des, src)
		if err != nil {
			return err
		}
		photoentity.ImageName = fileName
	}

	err := photorepository.UpdatePhotoById(photoentity)

	return err
}

func (s photoService) DeletePhoto(id string) error {

	photorepository := repository.NewPhotoRepository()
	oldphoto, err := photorepository.GetPhotoById(id)
	if err != nil {
		return err
	}
	if oldphoto.ImageName != "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		oldImageServerPath := filepath.Join(wd, "wwwroot", "images", "photocr", oldphoto.ImageName)
		os.Remove(oldImageServerPath)
	}

	err = photorepository.DeletePhotoById(id)

	return err
}

func (s photoService) AddVisitCount(id string) error {

	newsRepository := repository.NewPhotoRepository()
	photoo, err := newsRepository.GetPhotoById(id)
	if err != nil {
		return err
	}
	photoo.VisitCount += 1
	err = newsRepository.UpdatePhotoById(photoo)
	if err != nil {
		return err
	}

	return nil
}

func (s photoService) AddLike(id string) error {

	photoRepository := repository.NewPhotoRepository()
	photoo, err := photoRepository.GetPhotoById(id)
	if err != nil {
		return err
	}
	photoo.LikeCount += 1
	err = photoRepository.UpdatePhotoById(photoo)
	if err != nil {
		return err
	}

	return nil
}

func (s photoService) UploadPhoto(c echo.Context) error {
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
