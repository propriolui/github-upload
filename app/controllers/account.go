package controllers

import (
	"io/ioutil"
	"net/http"
	"propriolui/tracker_api/app/models"

	"go.uber.org/zap"
)

//Accounts : logger per accedere alla collezione account
type Accounts struct {
	accRepo models.AccountRepository
	s       *zap.SugaredLogger
}

//NewAccount : crea un nuovo oggetto account
func NewAccount(accRepo models.AccountRepository, s *zap.SugaredLogger) *Accounts {
	return &Accounts{accRepo, s}
}

//Login : permette di ricavare informazioni sull'account
func (a *Accounts) Login(w http.ResponseWriter, r *http.Request) {
	//legge il body, successivamente lo converte e lo passa al
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.s.Fatal(err)
	}

	bodyString := string(bodyBytes)
	result := a.accRepo.FindAccount(bodyString)
	a.s.Info(result)
}
