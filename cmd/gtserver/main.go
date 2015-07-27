package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jgroeneveld/gotrix/app/web"
	"github.com/jgroeneveld/gotrix/lib/db"
	"github.com/jgroeneveld/gotrix/lib/logger"
)

// TODO move to config object that reads from env etc
var (
	databaseURL    = "fixme"
	applicatioName = "fixme2"
)

func main() {
	l := logger.New()

	txManager, err := newTxManager()
	if err != nil {
		log.Fatal(err) // TODO output error stacktrace when available
	}

	router := web.NewRouter(l, txManager)

	port := getPort()
	l.Printf("Starting server on port=%s", port)

	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("Can not start server ERROR=%s", err.Error())
	}
}

func newTxManager() (*db.SimpleTxManager, error) {
	con, err := db.Connect(databaseURL, applicationName)
	if err != nil {
		return nil, err
	}
	return db.NewTxManager(db.NewTxFactory(con)), nil
}

// TODO move into config struct
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if !strings.Contains(port, ":") {
		port = "127.0.0.1:" + port
	}
	return port
}
