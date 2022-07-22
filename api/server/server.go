package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Service struct {
	items map[string]Item
	sync.RWMutex
}

func NewService(items map[string]Item) *Service {
	return &Service{
		items: items,
	}
}

func (s *Service) ListenAndServe() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/item", s.GetItems)
	r.Get("/item/{name}", s.GetItem)
	r.Post("/item", s.CreateItem)
	r.Put("/item/{name}", s.UpdateItem)
	r.Delete("/item/{name}", s.DeleteItem)

	log.Println("Starting server")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return err
	}
	return nil
}
