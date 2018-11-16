package controllers

import (
	"fmt"
	"net/http"

	"github.com/onkarbanerjee1/birdopedia/db"
	"github.com/onkarbanerjee1/birdopedia/models"
)

// BirdAdder handles POST request to add a bird
type BirdAdder struct {
	*db.DB
}

func (h *BirdAdder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	genericName := r.FormValue("name")
	if genericName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No name for bird? Please check"))

		return
	}
	birdBuilder := models.NewBirdBuilder(genericName)

	if commonName := r.FormValue("common_name"); commonName != "" {
		birdBuilder = birdBuilder.CommonName(commonName)
	}
	if scientificName := r.FormValue("scientific_name"); scientificName != "" {
		birdBuilder = birdBuilder.ScientificName(scientificName)
	}

	if err := h.Add(birdBuilder.Build()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Please try later, got %s", err)))

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Bird is Added !!"))
}
