package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Account : collezione account
type Account struct {
	ID               primitive.ObjectID  `bson:"_id, omitempty"`
	AccountID        string              `bson:"accountID"`
	AccountType      int8                `bson:"accountType"`
	RetainedEventAge int64               `bson:"retainedEventAge,omitempty"`
	MaximumDevices   int16               `bson:"maximumDevices,omitempty"`
	TotalPingCount   int16               `bson:"totalPingCount"`
	MaxPingCount     int8                `bson:"maxPingCount,omitempty"`
	AutoAddDevices   bool                `bson:"autoAddDevices,omitempty"`
	ExpirationTime   primitive.DateTime  `bson:"expirationTime"`
	SuspendUntilTime primitive.DateTime  `bson:"suspendUntilTime"`
	AllowWebService  bool                `bson:"allowWebService,omitempty"`
	AttributeMask    int64               `bson:"attributeMask,omitempty"`
	DefaultUser      string              `bson:"defaultUser,omitempty"`
	Password         string              `bson:"password"`
	TempPassword     string              `bson:"tempPassword,omitempty"`
	LastPassword     string              `bson:"lastPassword,omitempty"`
	PasswdChangeTime uint32              `bson:"passwdChangeTime,omitempty"`
	PasswdQueryTime  uint32              `bson:"passwdQueryTime,omitempty"`
	LastLoginTime    uint32              `bson:"lastLoginTime"`
	LoginMessage     string              `bson:"loginMessage,omitempty"`
	InactiveMessage  string              `bson:"inactiveMessage,omitempty"`
	DeletedTime      uint32              `bson:"deletedTime,omitempty"`
	IsActive         bool                `bson:"isActive"`
	Settings         *AccountSettings    `bson:"settings"`
	Info             *AccountInfo        `bson:"info"`
	Description      string              `bson:"description,omitempty"`
	Notes            string              `bson:"notes,omitempty"`
	LastUpdateTime   primitive.Timestamp `bson:"lastUpdateTime"`
	CreationTime     primitive.Timestamp `bson:"creationTime"`
}

//AccountSettings : settings dell'account
type AccountSettings struct {
	NotifyEmail      string `bson:"notifyEmail"`
	AllowNotify      bool   `bson:"allowNotify"`
	SpeedUnits       int    `bson:"speedUnits"`
	DistanceUnits    int    `bson:"distanceUnits"`
	VolumeUnits      int    `bson:"volumeUnits"`
	PressureUnits    int    `bson:"pressureUnits"`
	TemperatureUnits int    `bson:"temperatureUnits"`
	CurrencyUnits    int    `bson:"currencyUnits"`
	LatLonFormat     int    `bson:"latLonFormat"`
	GeocoderMode     int    `bson:"geocoderMode"`
	Timezone         string `bson:"timezone"`
	PreferDataFormat string `bson:"preferDataFormat"`
	PreferTimeFormat string `bson:"preferTimeFormat"`
}

//AccountInfo : informazioni account
type AccountInfo struct {
	PrivateName    bool   `bson:"privateName"`
	SMTPProperties string `bson:"smtpProperties,omitempty"`
	ContactName    string `bson:"contactName"`
	ContactPhone   string `bson:"contactPhone"`
	DisplayName    string `bson:"displayName,omitempty"`
}

//AccountRepository : funzioni da implementare che riguardano la collezione account
type AccountRepository interface {
	FindAccountByID(ID primitive.ObjectID) *Account
	FindAccount(AccountID string) *Account
	AddAccount(account *Account) interface{}
	UpdateAccount(account *Account)
}
