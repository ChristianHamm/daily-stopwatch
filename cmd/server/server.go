package main

import (
	"encoding/json"
	v1 "github.com/ChristianHamm/stopwatch/api/v1"
	"github.com/ChristianHamm/stopwatch/internal/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	// Handle static pages
	// Handle health
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// User API
	r.HandleFunc("/v1/user", v1.ListUsers).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/v1/user", v1.AddUser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/v1/user/{id:[0-9]+}", v1.ToggleUser).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/v1/user/{id:[0-9]+}", v1.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	r.PathPrefix("/").Handler(http.StripPrefix("/",
		http.FileServer(http.Dir("frontend/build"))))

	// Middleware
	r.Use(loggingMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(allowOriginMiddleware)

	// Main loop
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go updateSpeakDuration()
	log.Fatal(srv.ListenAndServe())
}

func updateSpeakDuration() {
	for range time.Tick(time.Second) {
		for i, user := range model.UserStore {
			if user.Speaking {
				model.UserStore[i].SpeakDuration++
			} else {
				model.UserStore[i].StartDate = time.Now()
			}
		}
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func allowOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
		next.ServeHTTP(w, req)
	})
}
