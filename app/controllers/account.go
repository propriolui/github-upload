package controllers

import (
	"encoding/json"
	"net/http"
	"propriolui/tracker_api/app/middlewares"
	"propriolui/tracker_api/app/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"go.uber.org/zap"
)

//Accounts : logger per accedere alla collezione account
type Accounts struct {
	accRepo models.AccountRepository
	s       *zap.SugaredLogger
}

type login struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

//NewAccount : crea un nuovo oggetto account
func NewAccount(accRepo models.AccountRepository, s *zap.SugaredLogger) *Accounts {
	return &Accounts{accRepo, s}
}

//Login : permette di ricavare informazioni sull'account
func (a *Accounts) Login(w http.ResponseWriter, r *http.Request) {

	msg := "Content-Type header is not application/json"
	err := middlewares.ValContentType(r)
	if err != nil {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	//legge il body, successivamente lo converte e lo passa al repo
	decoder := json.NewDecoder(r.Body)
	l := &models.Account{}
	err = decoder.Decode(&l)
	if err != nil {
		a.s.Panic(err)
	}
	a.s.Info(l)
	//recupera l'account associato all'indirizzo mail
	result := a.accRepo.FindAccount(l.AccountID)

	//se non esiste l'account esce altrimenti controlla la password
	if result.AccountID == "" {
		http.Error(w, "account not exist", http.StatusNotFound)
	} else {
		a.s.Info(l.Password)
		pwdRequest := []byte(l.Password)
		pwd := []byte(result.Password)

		// Comparing the password with the hash
		err = bcrypt.CompareHashAndPassword(pwd, pwdRequest)
		if err != nil {
			w.Header().Add("password", "incorrect")
		} else {
			w.Header().Add("password", "correct")
		}
	}
}

//CreateAccount : permette di creare un nuovo account
func (a *Accounts) CreateAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var err error

	acc := &models.Account{}
	acc.Info = &models.AccountInfo{}
	acc.Settings = &models.AccountSettings{}
	for key, values := range r.Form { // range over map
		for _, value := range values { // range over []string
			switch key {
			case "email":
				acc.AccountID = value
				break
			case "pwd":
				acc.Password = value
				break
			case "nEmail":
				acc.Settings.NotifyEmail = value
				break
			case "sUnits":
				acc.Settings.SpeedUnits, err = strconv.Atoi(value)
				if err != nil {
					a.s.Error(err)
				}
				break
			case "dUnits":
				acc.Settings.DistanceUnits, err = strconv.Atoi(value)
				if err != nil {
					a.s.Error(err)
				}
				break
			case "vUnits":
				acc.Settings.VolumeUnits, err = strconv.Atoi(value)
				if err != nil {
					a.s.Info(err)
				}
				break
			case "pUnits":
				acc.Settings.PressureUnits, err = strconv.Atoi(value)
				if err != nil {
					a.s.Error(err)
				}
				break
			case "tUnits":
				acc.Settings.TemperatureUnits, err = strconv.Atoi(value)
				if err != nil {
					a.s.Error(err)
				}
				break
			case "cUnits":
				acc.Settings.CurrencyUnits, err = strconv.Atoi(value)
				if err != nil {
					a.s.Error(err)
				}
				break
			case "lUnits":
				acc.Settings.LatLonFormat, err = strconv.Atoi(value)
				if err != nil {
					a.s.Error(err)
				}
				break
			case "timezone":
				acc.Settings.Timezone = value
				break
			case "dFormat":
				acc.Settings.PreferDataFormat = value
				break
			case "tformat":
				acc.Settings.PreferTimeFormat = value
				break
			case "pName":
				if value == "true" {
					acc.Info.PrivateName = true
				} else {
					acc.Info.PrivateName = false
				}
				break
			case "cName":
				acc.Info.ContactName = value
				break
			case "cPhone":
				acc.Info.ContactPhone = value
				break
			default:
				break
			}
		}
	}
	a.accRepo.AddAccount(acc)
}

func (a *Accounts) UpdateAccount(w http.ResponseWriter, r *http.Request) {

}
