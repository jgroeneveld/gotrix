package main

import (
	"log"
	"net/http"
	"strings"

	"gotrix/cfg"
	"gotrix/lib/db"
	"gotrix/lib/errors"
	"gotrix/lib/logger"
	"gotrix/web"
)

func main() {
	l := logger.New()

	txManager, err := newTxManager()
	if err != nil {
		log.Fatal(errors.ErrorWithStack(err))
	}

	router := web.NewRouter(l, txManager)

	port := getPort()
	l.Printf("Starting server on port=%s", port)

	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Can not start server\n%s", errors.ErrorWithStack(err))
	}
}

func newTxManager() (*db.SimpleTxManager, error) {
	con, err := db.Connect(cfg.Config.DatabaseURL, cfg.Config.ApplicationName+"_gtserver")
	if err != nil {
		return nil, err
	}
	return db.NewTxManager(db.NewTxFactory(con)), nil
}

// TODO move into config struct
func getPort() string {
	port := cfg.Config.Port
	if !strings.Contains(port, ":") {
		port = "127.0.0.1:" + port
	}
	return port
}
