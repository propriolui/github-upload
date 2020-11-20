package models

//User : collezione user
type User struct {
	AccountID         string `bson:"accountID"`
	UserID            string `bson:"userID"`
	UserType          int8   `bson:"userType"`
	RoleID            string `bson:"roleID"`
	Password          string `bson:"password"`
	TempPassword      string `bson:"tempPassword,omitempty"`
	LastPassword      string `bson:"lastPassword,omitempty"`
	Timezone          string `bson:"timezone"`
	FirstLoginPageID  string `bson:"firstLoginPageID,omitempty"`
	PreferredDeviceID string `bson:"preferredDeviceID,omitempty"`
	PasswdChangeTime  uint64 `bson:"passwdChangeTime,omitempty"`
	PasswdQueryTime   uint64 `bson:"passwdQueryTime,omitempty"`
	ExpirationTime    uint64 `bson:"expirationTime"`
	SuspendUntilTime  uint64 `bson:"suspendUntilTime,omitempty"`
	LastLoginTime     uint64 `bson:"lastLoginTime"`
	WelcomeTime       uint64 `bson:"welcomeTime"`
	IsActive          bool   `bson:"isActive"`
	DisplayName       string `bson:"displayName"`
	Description       string `bson:"description,omitempty"`
	Notes             string `bson:"notes,omitempty"`
	LastUpdateTime    uint64 `bson:"lastUpdateTime"`
	CreationTime      uint64 `bson:"creationTime"`
}

//UserAccessDetail : dettagli del livelo di accesso di un user
type UserAccessDetail struct {
	ACLID            string `bson:"aclID"`
	AccessLevel      uint8  `bson:"accessLevel"`
	LevelDescription string `bson:"levelDescription"`
}

//UserInfo : dettagli personali user
type UserInfo struct {
	NotifyEmail  string `bson:"notifyEmail"`
	ContactName  string `bson:"contactName"`
	ContactPhone string `bson:"contactPhone"`
	ContactEmail string `bson:"contactEmail"`
	Gender       int8   `bson:"gender"`
}
