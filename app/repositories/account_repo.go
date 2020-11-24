package repositories

import (
	"context"
	"propriolui/tracker_api/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

//AccountRepo implements models.AccountRepository
type AccountRepo struct {
	dba *mongo.Client
	s   *zap.SugaredLogger
}

//NewAccountRepo : crea una nuova repository per l'account
func NewAccountRepo(db *mongo.Client, s *zap.SugaredLogger) *AccountRepo {
	return &AccountRepo{db, s}
}

//FindAccount : ritorna un account in base all'accountID
func (acc *AccountRepo) FindAccount(aID string) *models.Account {
	collection := acc.dba.Database("tracker_db").Collection("account")
	filter := bson.D{{"accountID", aID}}
	result := &models.Account{}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		acc.s.Fatal(err)
	}
	acc.s.Info(result)
	return result
}

//AddAccount : ritorna un account in base all'accountID
func (acc *AccountRepo) AddAccount(account *models.Account) {
	return
}

//UpdateAccount : ritorna un account in base all'accountID
func (acc *AccountRepo) UpdateAccount(account *models.Account) {
	return
}
