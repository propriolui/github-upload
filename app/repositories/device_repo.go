package repositories

import (
	"context"
	"propriolui/tracker_api/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

//DeviceRepo : struct che rappresenta la classe
type DeviceRepo struct {
	dba *mongo.Client
	s   *zap.SugaredLogger
}

//NewDeviceRepo : instanzia la classe deviceRepo
func NewDeviceRepo(db *mongo.Client, s *zap.SugaredLogger) *DeviceRepo {
	return &DeviceRepo{db, s}
}

//CreateDevice : permette di creare un nuovo elemento nella tabella device
func (dev *DeviceRepo) CreateDevice(device *models.Device) interface{} {
	collection := dev.dba.Database("tracker_db").Collection("device")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	device.CreationTime = primitive.Timestamp{}
	insertResult, err := collection.InsertOne(ctx, device)
	if err != nil {
		dev.s.Panicf("error insert in db")
	}
	dev.s.Info("inserito nel db: ", insertResult.InsertedID)
	return insertResult.InsertedID
}

//UpdateDevice : modifica un device
func (dev *DeviceRepo) UpdateDevice(device *models.Device) {
	collection := dev.dba.Database("tracker_db").Collection("device")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	device.LastUpdateTime = primitive.Timestamp{}
	result, err := collection.ReplaceOne(ctx, bson.M{"_id": device.ID}, device)
	if err != nil {
		dev.s.DPanic(err)
	}
	dev.s.Info("Modificato: ", result.ModifiedCount)
}

//FindDevice : ritorna un device in base all'ID
func (dev *DeviceRepo) FindDevice(ID primitive.ObjectID) *models.Device {
	collection := dev.dba.Database("tracker_db").Collection("device")
	filter := bson.M{"_id": ID}
	result := &models.Device{}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			dev.s.Infof("documento vuoto")
		} else {
			dev.s.Panic(err)
		}
	}
	return result
}

/*
//FindAllAccountDevices : ritorna tutti i device di un determinato account
func (dev *DeviceRepo) FindAllAccountDevices(accountID primitive.ObjectID) []models.Device {
	collection := dev.dba.Database("tracker_db").Collection("device")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var result []models.Device

}*/
