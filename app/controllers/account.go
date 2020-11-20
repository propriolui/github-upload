package controllers

import (
	"net/http"
	"propriolui/tracker_api/app/models"
)

//Accounts : logger per accedere alla collezione account
type Accounts struct {
	accRepo models.AccountRepository
}

//NewAccount : crea un nuovo oggetto account
func NewAccount(accRepo models.AccountRepository) *Accounts {
	return &Accounts{accRepo}
}

//GetAccount : permette di ricavare informazioni sull'account
func (a *Accounts) Login(w http.ResponseWriter, r *http.Request) {

}
