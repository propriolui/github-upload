package models

//StatusCode : collezione StatusCode
type StatusCode struct {
	AccountID       string `bson:"accountID"`
	DeviceID        string `bson:"deviceID"`
	StatusCode      uint32 `bson:"statusCode"`
	StatusName      string `bson:"statusName"`
	ForegroundColor string `bson:"foregroundColor"`
	BackgroundColor string `bson:"backgroundColor"`
	IconSelector    string `bson:"iconSelector"`
	IconName        string `bson:"iconName"`
	Description     string `bson:"description,omitempty"`
	LastUpdateTime  uint32 `bson:"lastUpdateTime"`
	CreationTime    uint32 `bson:"creationTime"`
}
