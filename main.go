package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/onkarbanerjee1/birdopedia/birds"
	cfg "github.com/onkarbanerjee1/birdopedia/config"
)

func main() {

	env, err := cfg.NewEnv("postgres://onkar:passwd@localhost/bird_db")
	if err != nil {
		panic(err)
	}
	defer env.DB.Close()

	r := mux.NewRouter()

	r.Handle("/birds/query", birds.QueryPage()).Methods(http.MethodGet)
	r.Handle("/birds/fetch", birds.GetBirdByName(env)).Methods(http.MethodGet)
	r.Handle("/birds/add", birds.NewBirdForm()).Methods(http.MethodGet)
	r.Handle("/birds/add", birds.InsertNewBird(env)).Methods(http.MethodPost)
	r.Handle("/birds/update", birds.UpdateBirdForm()).Methods(http.MethodGet)
	r.Handle("/birds/update", birds.UpdateBird(env)).Methods(http.MethodPost)
	r.Handle("/birds/delete", birds.DeleteBirdForm()).Methods(http.MethodGet)
	r.Handle("/birds/delete", birds.DeleteBird(env)).Methods(http.MethodPost)
	r.Handle("/birds", birds.GetAllBirds(env)).Methods(http.MethodGet)
	r.Handle("/birds/{name}", birds.GetBirdsByName(env)).Methods(http.MethodGet)
	r.Handle("/", birds.MainPage()).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: r,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
