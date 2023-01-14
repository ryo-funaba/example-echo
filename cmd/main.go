package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	adaptorHTTP "github.com/ryo-funaba/example_echo/internal/adapter/http"
)

func main() {
	router := adaptorHTTP.InitRouter()
	addr := os.Getenv("HTTP_PORT")

	rTimeout, err := strconv.Atoi(os.Getenv("HTTP_ReadTimeout"))
	if err != nil {
		panic(err)
	}

	wTimeout, err := strconv.Atoi(os.Getenv("HTTP_WriteTimeout"))
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:         ":" + addr,
		Handler:      router,
		ReadTimeout:  time.Duration(rTimeout) * time.Second,
		WriteTimeout: time.Duration(wTimeout) * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}
