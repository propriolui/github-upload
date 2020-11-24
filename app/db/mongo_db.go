package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

//Mongodb : struct per creare un nuovo db
type Mongodb struct {
	s *zap.SugaredLogger
}

//NewDB : crea un nuovo oggetto db
func NewDB(s *zap.SugaredLogger) *Mongodb {
	return &Mongodb{s}
}

//ConnectDB : apre la connessione con il db mongo
func (mdb *Mongodb) ConnectDB() (*mongo.Client, context.Context) {

	//apre una nuova connessione col db
	clientDB, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tracking-web-app:6OW0KwrITghWTYJW@gabryelcluster.xdyuo.mongodb.net/tracker_db?retryWrites=true&w=majority"))
	if err != nil {
		mdb.s.Error(err)
	}
	//aspetta il collegamento del db per 10 secondi, in caso da' errore
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = clientDB.Connect(ctx)
	if err != nil {
		mdb.s.Error(err)
	}
	defer cancel()

	//controllo di connessione riuscita
	err = clientDB.Ping(ctx, readpref.Primary())
	if err != nil {
		mdb.s.Error(err)
	}

	//controllo di connessione riuscita
	err = clientDB.Ping(ctx, readpref.Primary())
	if err != nil {
		mdb.s.Error(err)
	}

	return clientDB, ctx
}

//Disconnect : disconnetti dal db
func (mdb *Mongodb) Disconnect(ctx context.Context, db *mongo.Client) {
	db.Disconnect(ctx)
}
