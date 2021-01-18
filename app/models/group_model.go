package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Group : collezione group
type Group struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	AccountID      string             `bson:"accountID"`
	GroupID        string             `bson:"groupID"`
	DeviceID       string             `bson:"deviceID"`
	DisplayName    string             `bson:"displayName"`
	Notes          string             `bson:"notes, omitempty"`
	LastUpdateTime uint64             `bson:"lastUpdateTime"`
	CreationTime   uint64             `bson:"creationTime"`
}
