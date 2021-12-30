package main

import (
	"log"
	"net/http"
	"time"

	"github.com/zhimiaox/gin-swagger/route"

	"golang.org/x/sync/errgroup"

	_ "github.com/zhimiaox/gin-swagger/docs/swa"
	_ "github.com/zhimiaox/gin-swagger/docs/swb"
)

var (
	g errgroup.Group
)

//go:generate swag init --instanceName=swa --generalInfo=route/swa.go --exclude=api/swb --output docs/swa
//go:generate swag init --instanceName=swb --generalInfo=route/swb.go --exclude=api/swa --output docs/swb

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      route.Swa(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      route.Swb(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
