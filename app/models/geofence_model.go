package models

//Geofence : collezione geofence
type Geofence struct {
	AccountID           string  `bson:"accountID"`
	GeofenceID          string  `bson:"geofenceID"`
	MinLatitude         float64 `bson:"minLatitude"`
	MaxLatitude         float64 `bson:"maxLatitude"`
	MinLongitude        float64 `bson:"minLongitude"`
	MaxLongitude        float64 `bson:"maxLongitude"`
	ZonePurposeID       string  `bson:"zonePurposeID,omitempty"`
	ReverseGeocode      bool    `bson:"reverseGeocode"`
	ArrivalZone         bool    `bson:"arrivalZone,omitempty"`
	ArrivalStatusCode   uint16  `bson:"arrivalStatusCode,omitempty"`
	DepartureZone       bool    `bson:"despartureZone,omitempty"`
	DepartureStatusCode uint16  `bson:"departureStatusCode,omitempty"`
	AutoNotify          bool    `bson:"autoNotify"`
	ZoomRegion          bool    `bson:"zoomRegion"`
	ShapeColor          string  `bson:"shapeColor"`
	IconName            string  `bson:"iconName"`
	ZoneType            uint8   `bson:"zoneType,omitempty"`
	Radius              uint16  `bson:"radius"`
	Vertices            string  `bson:"vertices,omitempty"`
	Latitude1           float32 `bson:"latitude1"`
	Longitude1          float32 `bson:"longitude1"`
	Latitude2           float32 `bson:"latitude2"`
	Longitude2          float32 `bson:"longitude2"`
	Latitude3           float32 `bson:"latitude3"`
	Longitude3          float32 `bson:"longitude3"`
	Latitude4           float32 `bson:"latitude4"`
	Longitude4          float32 `bson:"longitude4"`
	Latitude5           float32 `bson:"latitude5"`
	Longitude5          float32 `bson:"longitude5"`
	Latitude6           float32 `bson:"latitude6"`
	Longitude6          float32 `bson:"longitude6"`
	Latitude7           float32 `bson:"latitude7"`
	Longitude7          float32 `bson:"longitude7"`
	Latitude8           float32 `bson:"latitude8"`
	Longitude8          float32 `bson:"longitude8"`
	Latitude9           float32 `bson:"latitude9"`
	Longitude9          float32 `bson:"longitude9"`
	Latitude10          float32 `bson:"latitude10"`
	Longitude10         float32 `bson:"longitude10"`
	GroupID             string  `bson:"groupID"`
	StreetAddress       string  `bson:"streetAddress,omitempty"`
	City                string  `bson:"city,omitempty"`
	StateProvince       string  `bson:"stateProvince,omitempty"`
	PostalCode          string  `bson:"postalCode,omitempty"`
	Country             string  `bson:"country,omitempty"`
	Subdivision         string  `bson:"subdivision,omitempty"`
	IsActive            bool    `bson:"isActive"`
	DisplayName         string  `bson:"displayName"`
	Description         string  `bson:"descriprion,omitempty"`
	LastUpdateTime      uint32  `bson:"lastUpdateTime"`
	CreationTime        uint32  `bson:"creationTime"`
}
