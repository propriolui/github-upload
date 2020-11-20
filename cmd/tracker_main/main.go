package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"propriolui/tracker_api/app/controllers"
	"propriolui/tracker_api/app/db"
	"propriolui/tracker_api/app/repositories"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	//mini main
	run()
}

func run() {

	//creazione logger
	logger, _ := zap.NewProduction()

	//chiusura logger
	defer logger.Sync()

	//creo un sugar logger
	sugar := logger.Sugar()

	//creo un nuovo oggetto database
	dbmongo := db.NewDB(sugar)

	//connessione al db
	clientDB, ctx := dbmongo.ConnectDB()

	//disconnette dal database quando finisce l'esecuzione
	defer dbmongo.Disconnect(ctx, clientDB)

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	//crea repositories
	accRepo := repositories.NewAccountRepo(clientDB)

	a := controllers.NewAccount(accRepo)

	r := customRouter()

	// crea un nuovo server
	s := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      ch(r),             // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start server
	go func() {
		sugar.Info("Start server on port 8080")
		err := s.ListenAndServe()
		if err != nil {
			sugar.Error(err)
			os.Exit(1)
		}
	}()

	// spegnimento dolce server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Blocca il processo fino a quando non arriva un segnale
	sig := <-c
	sugar.Infof("Got signal:", sig)

	// spegne dolcemente il server ed aspetta un massimo di 30 secondi per far completare tutte le operazioni in corso
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}

func customRouter() *mux.Router {
	r := mux.NewRouter()

	return r
}
