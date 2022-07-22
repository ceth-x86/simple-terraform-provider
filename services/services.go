package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Entity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Server struct {
	Entities map[string]Entity
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (s *Server) GetEntities(w http.ResponseWriter, r *http.Request) {
	var entities []Entity
	for _, entity := range s.Entities {
		entities = append(entities, entity)
	}
	respondwithJSON(w, http.StatusOK, entities)
}

func (s *Server) CreateEntity(w http.ResponseWriter, r *http.Request) {
	var entity Entity
	json.NewDecoder(r.Body).Decode(&entity)

	s.Entities[entity.Id] = entity
	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func (s *Server) UpdateEntity(w http.ResponseWriter, r *http.Request) {
	var entity Entity
	json.NewDecoder(r.Body).Decode(&entity)
	id := chi.URLParam(r, "id")

	s.Entities[id] = entity
	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully updated"})
}

func (s *Server) DeleteEntity(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	delete(s.Entities, id)
	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func (s *Server) GetEntity(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	entity := s.Entities[id]
	respondwithJSON(w, http.StatusOK, entity)
}

func main() {
	server := &Server{
		Entities: make(map[string]Entity),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/entities", server.GetEntities)
	r.Get("/entities/{id}", server.GetEntity)
	r.Post("/entities", server.CreateEntity)
	r.Put("/entities/{id}", server.UpdateEntity)
	r.Delete("/entities/{id}", server.DeleteEntity)
	http.ListenAndServe(":3000", r)
}
