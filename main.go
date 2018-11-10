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
	"github.com/onkarbanerjee1/birdopedia/bird"
)

var db *bird.DB

func main() {

	db = bird.NewDB()

	hawkBuilder := bird.NewBuilder("Hawk")
	hawk := hawkBuilder.CommonName("Northern goshawk").ScientificName("Accipiter gentilis").Habitat([]string{"Asia", "Europe",
		"North America"}).Endangered(true).PostedBy("Onkar Banerjee").Build()
	db.Add(hawk)

	eagleBuilder := bird.NewBuilder("Eagle")
	eagle := eagleBuilder.CommonName("Harpy eagle").ScientificName("Harpia harpyja").Habitat([]string{"South America",
		"Central America"}).Endangered(true).PostedBy("Tuhina Banerjee").Build()
	db.Add(eagle)

	fmt.Println("Contents of db are \n", db)

	r := mux.NewRouter()

	r.HandleFunc("/birds", getAllBirds).Methods(http.MethodGet)
	r.HandleFunc("/birds/{name}", getBirdByGenericName).Methods(http.MethodGet)
	r.HandleFunc("/birds", addBird).Methods(http.MethodPost)

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

func getAllBirds(w http.ResponseWriter, r *http.Request) {
	birds := db.GetAll()

	w.WriteHeader(http.StatusOK)

	for _, each := range birds {
		w.Write([]byte(each.String()))
	}
}

func getBirdByGenericName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bird, err := db.GetByGenericName(vars["name"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(bird.String()))

}

func addBird(w http.ResponseWriter, r *http.Request) {
	genericName := r.FormValue("name")
	if genericName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No name for bird? Please check"))

		return
	}
	birdBuilder := bird.NewBuilder(genericName)

	if commonName := r.FormValue("common_name"); commonName != "" {
		birdBuilder = birdBuilder.CommonName(commonName)
	}
	if scientificName := r.FormValue("scientific_name"); scientificName != "" {
		birdBuilder = birdBuilder.ScientificName(scientificName)
	}

	if err := db.Add(birdBuilder.Build()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Please try later, got %s", err)))

		return
	}

	w.WriteHeader(http.StatusOK)
}
