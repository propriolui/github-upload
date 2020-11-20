package models

//Device : collezione device
type Device struct {
	AccountID          string  `bson:"accountID"`
	DeviceID           string  `bson:"deviceID"`
	GroupID            string  `bson:"groupID"`
	EquipmentType      string  `bson:"equipmentType,omitempty"`
	EquipmentStatus    string  `bson:"equipmentStatus,omitempty"`
	SpeedLimitKMH      float32 `bson:"speedLimitKMH,omitempty"`
	SpeedOffsetKMH     float32 `bson:"speedOffsetKMH,omitempty"`
	PlanDistanceKM     float32 `bson:"planDistanceKM,omitempty"`
	InstallTime        uint32  `bson:"installTime,omitempty"`
	ResetTime          uint32  `bson:"resetTime,omitempty"`
	ExpirationTime     uint32  `bson:"expirationTime,omitempty"`
	DeviceCode         string  `bson:"deviceCode,omitempty"`
	IconURL            string  `bson:"iconURL,omitempty"`
	SerialNumber       string  `bson:"serialNumber"`
	SimPhoneNumber     string  `bson:"simPhoneNumber"`
	SimID              string  `bson:"simID,omitempty"`
	ImeiNumber         string  `bson:"imeiNumber,omitempty"`
	RouteColor         string  `bson:"routeColor,omitempty"`
	IgnitionIndex      uint8   `bson:"ignitionIndex,omitempty"`
	CodeVersion        string  `bson:"codeVersion,omitempty"`
	FeatureSet         string  `bson:"featureSet,omitempty"`
	IPAddressValid     string  `bson:"ipAddressValid"`
	FirstConnectTime   uint32  `bson:"firstConnectTime"`
	LastRevGeocodeTime uint32  `bson:"lastReverseGeocodeTime,omitempty"`
	CommandStateMask   uint32  `bson:"commandStateMask,omitempty"`
	StatusCodeState    uint32  `bson:"statusCodeState"` //on\off state
	OdometerOffsetKM   float32 `bson:"odometerOffsetKM,omitempty"`
	LastStopTime       uint32  `bson:"lastStopTime,omitempty"`
	LastStartTime      uint32  `bson:"lastStartTime,omitempty"`
	RegistrationTime   uint32  `bson:"registrationTime"`
	DeletedTime        uint32  `bson:"deletedTime"`
	IsActive           bool    `bson:"isActive"`
	DisplayName        string  `bson:"displayName"`
	Notes              string  `bson:"notes,omitempty"`
	LastUpdateTime     uint32  `bson:"lastUpdateTime"`
	CreationTime       uint32  `bson:"creationTime"`
}

//LastInfo : informazioni su tutti i parametri last
type LastInfo struct {
	LastInputState      uint32  `bson:"lastInputState,omitempty"`
	LastOutputState     uint32  `bson:"lastOutputState,omitempty"`
	LastBatteryLevel    float64 `bson:"lastBatteryLevel"`
	LastBatteryVolts    float64 `bson:"lastBatteryVolts,omitempty"`
	LastValidLatitude   float64 `bson:"lastValidLatitude"`
	LastValidLongitude  float64 `bson:"lastValidLongitude"`
	LastValidHeading    float64 `bson:"lastValidHeading"`
	LastValidSpeedKMH   float64 `bson:"lastValidSpeedKMH,omitempty"`
	LastGPSTimestamp    uint32  `bson:"lastGPSTimestamp"`
	LastEventTimestamp  uint32  `bson:"lastEventTimestamp"`
	LastEventStatusCode uint32  `bson:"lastEventStatusCode"`
	LastCellServingInfo string  `bson:"lastCellServingInfo,omitempty"`
	LastDistanceKM      float32 `bson:"lastDistanceKM,omitempty"`
	LastOdometerKM      float32 `bson:"lastOdometerKM,omitempty"`
}

//ConnectionInfo : informazioni sulla connessione TCP
type ConnectionInfo struct {
	ExpextAck             bool   `bson:"expectAck"`
	LastAckCommand        string `bson:"lastAckCommand"`
	LastAckTime           uint32 `bson:"lastAckTime"`
	FixedTCPSessionID     string `bson:"fixedTcpSessionID,omitempty"`
	LastTCPSessionID      string `bson:"lastTCPSessionID,omitempty"`
	IPAddressCurrent      string `bson:"ipAddressCurrent"`
	RemotePortCurrent     uint16 `bson:"remotePortCurrent"`
	ListenPortCurrent     uint16 `bson:"listenPostCurrent"`
	LastTotalConnectTime  uint32 `bson:"lastTotalConnectTime,omitempty"`
	LastDuplexConnectTime uint32 `bson:"lastDuplexConnectTime,omitempty"`
	PendingPingCommand    string `bson:"pendingPingCommand,omitempty"`
	LastPingTime          uint32 `bson:"lastPingTime,omitempty"`
	TotalPingCount        uint32 `bson:"totalPingCount,omitempty"`
	MaxPingCount          uint32 `bson:"maxPingCount,omitempty"`
}

//DCSInfo : informazioni sui codici di tracking {opzionale}
type DCSInfo struct {
	DCSPropertiesID string `bson:"dcsPropertiesID,omitempty"`
	DCSConfigMask   uint32 `bson:"dcsConfigMask,omitempty"`
	DCSConfigString string `bson:"dcsConfigString,omitempty"`
	DCSCommandHost  string `bson:"dcsCommandHost,omitempty"`
	DCSCommandState string `bson:"dcsCommandState,omitempty"`
}

//VehicleInfo : informazioni sul veicolo {opzionale}
type VehicleInfo struct {
	VehicleMake  string `bson:"vehicleMake,omitempty"`
	VehicleModel string `bson:"vehicleModel,omitempty"`
	VehicleColor string `bson:"vehicleColor,omitempty"`
	VehicleYear  uint16 `bson:"vehicleYear,omitempty"`
	VehicleID    string `bson:"vehicleID,omitempty"`
}
