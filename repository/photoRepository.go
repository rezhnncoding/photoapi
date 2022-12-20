package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"puppy/database"
	"puppy/model/photo"
)

type PhotoRepository interface {
	GetPhotoList() ([]photo.Photo, error)
	GetPhotoById(id string) (photo.Photo, error)
	InsertPhoto(user photo.Photo) (string, error)
	UpdatePhotoById(user photo.Photo) error
	DeletePhotoById(id string) error
}

type photoRepository struct {
	db database.Db
}

func NewPhotoRepository() PhotoRepository {
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	return photoRepository{
		db: db,
	}
}

func (photoRepository photoRepository) GetPhotoList() ([]photo.Photo, error) {

	photoCollection := photoRepository.db.GetNewsCollection()

	cursor, err := photoCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var news []photo.Photo
	err = cursor.All(context.TODO(), &news)
	if err != nil {
		return nil, err
	}

	return news, nil

}

func (photoRepository photoRepository) GetPhotoById(id string) (photo.Photo, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return photo.Photo{}, err
	}
	photoCollection := photoRepository.db.GetNewsCollection()
	var userObject photo.Photo
	// db.getCollection('users').find({"_id" : ObjectId("6297bdcc8d7757574658ed66")})
	err = photoCollection.FindOne(context.TODO(), bson.D{
		{"_id", objectId},
	}).Decode(&userObject)

	if err != nil {
		return photo.Photo{}, err
	}

	return userObject, nil

}

func (photoRepository photoRepository) InsertPhoto(photo photo.Photo) (string, error) {

	photoCollection := photoRepository.db.GetNewsCollection()
	res, err := photoCollection.InsertOne(context.TODO(), photo)

	if err != nil {
		return "", err
	}
	objectId := res.InsertedID.(primitive.ObjectID).Hex()
	return objectId, nil
}

func (photoRepository photoRepository) UpdatePhotoById(photo photo.Photo) error {
	objectId, err := primitive.ObjectIDFromHex(photo.Id)
	if err != nil {
		return err
	}
	photo.Id = ""
	photoCollection := photoRepository.db.GetNewsCollection()
	_, err = photoCollection.UpdateOne(context.TODO(), bson.D{{"_id", objectId}}, bson.D{{"$set", photo}})

	if err != nil {
		return err
	}

	return nil
}

func (photoRepository photoRepository) DeletePhotoById(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	photoCollection := photoRepository.db.GetNewsCollection()
	_, err = photoCollection.DeleteOne(context.TODO(), bson.D{{"_id", objectId}})

	if err != nil {
		return err
	}

	return nil
}
