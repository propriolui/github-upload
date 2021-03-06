package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//StatusCode : collezione StatusCode
type StatusCode struct {
	ID              primitive.ObjectID `bson:"_id, omitempty"`
	AccountID       primitive.ObjectID `bson:"accountID"`
	DeviceID        string             `bson:"deviceID"`
	StatusCode      uint32             `bson:"statusCode"`
	StatusName      string             `bson:"statusName"`
	ForegroundColor string             `bson:"foregroundColor"`
	BackgroundColor string             `bson:"backgroundColor"`
	IconSelector    string             `bson:"iconSelector"`
	IconName        string             `bson:"iconName"`
	Description     string             `bson:"description,omitempty"`
	LastUpdateTime  uint32             `bson:"lastUpdateTime"`
	CreationTime    uint32             `bson:"creationTime"`
}
