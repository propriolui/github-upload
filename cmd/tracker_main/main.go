package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"propriolui/tracker_api/app/controllers"
	"propriolui/tracker_api/app/db"
	"propriolui/tracker_api/app/middlewares"
	"propriolui/tracker_api/app/repositories"
	"time"

	"github.com/rs/cors"

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
	ch := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	//crea repositories
	accRepo := repositories.NewAccountRepo(clientDB, sugar)

	a := controllers.NewAccount(accRepo, sugar)

	r := mux.NewRouter()

	handler := ch.Handler(r)

	// crea un nuovo server
	s := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      handler,           // set the default handler
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

	//routing
	r.HandleFunc("/login", a.Login).Methods("POST")
	r.HandleFunc("/accountByID", middlewares.IsAuthorized(a.GetAccountByToken)).Methods("GET")
	r.HandleFunc("/account", middlewares.IsAuthorized(a.GetAccount)).Methods("GET")
	r.HandleFunc("/account", a.CreateAccount).Methods("POST")
	r.HandleFunc("/account", middlewares.IsAuthorized(a.UpdateAccount)).Methods("PUT")

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
