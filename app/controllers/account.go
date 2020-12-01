package controllers

import (
	"encoding/json"
	"net/http"
	"propriolui/tracker_api/app/models"

	"github.com/golang/gddo/httputil/header"

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

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	//legge il body, successivamente lo converte e lo passa al
	decoder := json.NewDecoder(r.Body)
	var l login
	err := decoder.Decode(&l)
	if err != nil {
		panic(err)
	}
	a.s.Info(l)
	result := a.accRepo.FindAccount(l.Email)
	if result.AccountID == "" {
		http.Error(w, "account not exist", http.StatusNotFound)
	} else {
		if result.Password == l.Password {
			w.Header().Add("password", "correct")
		} else {
			http.Error(w, "password not matching", http.StatusNotAcceptable)
		}
	}
}
