package repositories

import (
	"context"
	"propriolui/tracker_api/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		if err == mongo.ErrNoDocuments {
			acc.s.Infof("documento vuoto")
		} else {
			acc.s.Panic(err)
		}
	}
	return result
}

//AddAccount : inserisce un account nel db
func (acc *AccountRepo) AddAccount(account *models.Account) {
	collection := acc.dba.Database("tracker_db").Collection("account")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	account.CreationTime = primitive.Timestamp{}
	insertResult, err := collection.InsertOne(ctx, account)
	if err != nil {
		acc.s.Panicf("error insert in db")
	}
	acc.s.Info("inserito nel db: ", insertResult.InsertedID)
	return
}

//UpdateAccount : ritorna un account in base all'accountID
func (acc *AccountRepo) UpdateAccount(account *models.Account) {
	collection := acc.dba.Database("tracker_db").Collection("account")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	account.LastUpdateTime = primitive.Timestamp{}
	result, err := collection.ReplaceOne(ctx, bson.M{"accountID": account.AccountID}, account)
	if err != nil {
		acc.s.DPanic(err)
	}
	acc.s.Info("Modificato: ", result.ModifiedCount)
}
