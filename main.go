package main

import (
	"net/http"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		var stopt storage.Client
		var psopt pubsub.Client
		logrus.Infof("request %s, %v, %v", uuid.NewString(), stopt, psopt)
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
