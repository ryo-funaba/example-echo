package main

import (
	"log"
	"net/http"

	"github.com/ryo-funaba/example_echo/internal/config"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	s := &http.Server{
		Addr:         conf.HTTPInfo.Addr,
		ReadTimeout:  conf.HTTPInfo.ReadTimeout,
		WriteTimeout: conf.HTTPInfo.WriteTimeout,
	}

	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}
