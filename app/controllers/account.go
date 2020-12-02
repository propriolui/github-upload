package controllers

import (
	"encoding/json"
	"net/http"
	"propriolui/tracker_api/app/models"

	"github.com/golang/gddo/httputil/header"
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
		a.s.Panic(err)
	}
	a.s.Info(l)
	//recupera l'account associato all'indirizzo mail
	result := a.accRepo.FindAccount(l.Email)

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
