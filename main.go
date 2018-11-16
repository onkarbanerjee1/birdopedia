package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/onkarbanerjee1/birdopedia/controllers"
	"github.com/onkarbanerjee1/birdopedia/db"
	"github.com/onkarbanerjee1/birdopedia/models"
)

func main() {

	db := db.NewDB()

	hawkBuilder := models.NewBirdBuilder("Hawk")
	hawk := hawkBuilder.CommonName("Northern goshawk").ScientificName("Accipiter gentilis").Habitat([]string{"Asia", "Europe",
		"North America"}).Endangered(true).PostedBy("Onkar Banerjee").Build()
	db.Add(hawk)

	eagleBuilder := models.NewBirdBuilder("Eagle")
	eagle := eagleBuilder.CommonName("Harpy eagle").ScientificName("Harpia harpyja").Habitat([]string{"South America",
		"Central America"}).Endangered(true).PostedBy("Tuhina Banerjee").Build()
	db.Add(eagle)

	fmt.Println("Contents of db are \n", db)

	r := mux.NewRouter()

	r.HandleFunc("/form", controllers.FormHandler)
	r.Handle("/birds", &controllers.AllBirdsGetter{DB: db}).Methods(http.MethodGet)
	r.Handle("/birds/{name}", &controllers.BirdGetter{DB: db}).Methods(http.MethodGet)
	r.Handle("/addBirds", &controllers.BirdAdder{DB: db}).Methods(http.MethodPost)

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
