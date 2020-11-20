package repositories

import (
	"propriolui/tracker_api/app/db"
	"propriolui/tracker_api/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//AccountRepo : struct addetta ad accede alle funzioni di Account
type AccountRepo struct {
	dba *mongo.Client
}

//NewAccountRepo : crea una nuova repository per l'account
func NewAccountRepo(db *mongo.Client) *AccountRepo {
	return &AccountRepo{db}
}

//FindAccount : ritorna un account in base all'accountID
func (acc *AccountRepo) FindAccount(a string) (*models.Account, error){
	collection := acc.dba.Database().Collection("account")
	var account bson.M
	if err = collection.FindOne(ctx, ) {
		
	}
}

//AddAccount : ritorna un account in base all'accountID
func AddAccount(a string) {
}

//UpdateAccount : ritorna un account in base all'accountID
func UpdateAccount(a string) {
}
