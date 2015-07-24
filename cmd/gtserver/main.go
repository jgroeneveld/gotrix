package main

import (
	"github.com/jgroeneveld/gotrix/app/web"
	"github.com/jgroeneveld/gotrix/lib/logger"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	l := logger.New()
	router := web.NewRouter(l)

	port := getPort()
	l.Printf("Starting server on port=%s", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("Can not start server ERROR=%s", err.Error())
	}
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
