package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sfardiansyah/laywook/pkg/auth"
)

// Handler ...
func Handler(a auth.Service) http.Handler {
	r := mux.NewRouter()

	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/users", getUser(a)).Methods("GET")

	return r
}

func getUser(a auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		lR := GetUserRequest{}
		if err := json.NewDecoder(r.Body).Decode(&lR); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := a.GetUser(lR.Email)
		if err != nil {
			if errors.Is(err, auth.ErrInvalidCredentials) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payload := map[string]interface{}{
			"user": user,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	}
}
