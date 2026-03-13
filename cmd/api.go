package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	// logger
	// db driver
}

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all good"))
	})

	
	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: h,
		ReadTimeout: 10 * time.Second, 
		WriteTimeout: 30 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

	log.Printf("server has started at addr %s", app.config.addr)
	return srv.ListenAndServe()
}

type config struct {
	addr string "8080"
	db   dbConfig
}

type dbConfig struct {
	dsn string // domain string

}