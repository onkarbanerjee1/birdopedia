package controllers

import (
	"net/http"

	"github.com/onkarbanerjee1/birdopedia/db"
)

// AllBirdsGetter handles GET request for getting all birds
type AllBirdsGetter struct {
	*db.DB
}

func (h *AllBirdsGetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	birds := h.GetAll()

	w.WriteHeader(http.StatusOK)

	for _, each := range birds {
		w.Write([]byte(each.String()))
	}
}
