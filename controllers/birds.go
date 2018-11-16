package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onkarbanerjee1/birdopedia/db"
)

// BirdGetter handles GET request for getting birds by generic name
type BirdGetter struct {
	*db.DB
}

func (h *BirdGetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bird, err := h.GetByGenericName(vars["name"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(bird.String()))
}
