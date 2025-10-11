package utils

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetIDParam(w http.ResponseWriter, r *http.Request, paramName string) (int, bool) {
	idStr := chi.URLParam(r, paramName)
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid "+paramName, http.StatusBadRequest)
		return 0, false
	}

	return id, true
}