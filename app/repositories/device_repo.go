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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": ID}
	result := &models.Device{}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			dev.s.Infof("documento vuoto")
		} else {
			dev.s.Panic(err)
		}
	}
	return result
}

//FindAllAccountDevices : ritorna tutti i device di un determinato account
func (dev *DeviceRepo) FindAllAccountDevices(accountID primitive.ObjectID) []models.Device {
	collection := dev.dba.Database("tracker_db").Collection("device")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var result []models.Device
	cursor, err := collection.Find(ctx, bson.M{"accountID": accountID})
	if err != nil {
		dev.s.DPanic(err)
	}
	defer cursor.Close(ctx)
	i := 0
	for cursor.Next(ctx) {
		if err = cursor.Decode(&result[i]); err != nil {
			dev.s.Fatal(err)
		}
		dev.s.Info(result[i])
	}
	return result
}
