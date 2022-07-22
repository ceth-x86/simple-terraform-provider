package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *Service) GetItems(w http.ResponseWriter, r *http.Request) {
	s.RLock()
	defer s.RUnlock()
	err := json.NewEncoder(w).Encode(s.items)
	if err != nil {
		log.Println(err)
	}
}

func (s *Service) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	s.Lock()
	defer s.Unlock()

	if s.itemExists(item.Name) {
		http.Error(w, fmt.Sprintf("item %s already exists", item.Name), http.StatusBadRequest)
		return
	}

	s.items[item.Name] = item
	log.Printf("added item: %s", item.Name)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Printf("error sending response - %s", err)
	}
}

func (s *Service) UpdateItem(w http.ResponseWriter, r *http.Request) {
	itemName := chi.URLParam(r, "name")
	if itemName == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	var item Item
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	s.Lock()
	defer s.Unlock()

	if !s.itemExists(itemName) {
		log.Printf("item %s does not exist", itemName)
		http.Error(w, fmt.Sprintf("item %v does not exist", itemName), http.StatusBadRequest)
		return
	}

	s.items[itemName] = item
	log.Printf("updated item: %s", item.Name)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Printf("error sending response - %s", err)
	}
}

func (s *Service) DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemName := chi.URLParam(r, "name")
	if itemName == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	s.Lock()
	defer s.Unlock()

	if !s.itemExists(itemName) {
		http.Error(w, fmt.Sprintf("item %s does not exists", itemName), http.StatusNotFound)
		return
	}

	delete(s.items, itemName)

	_, err := fmt.Fprintf(w, "Deleted item with name %s", itemName)
	if err != nil {
		log.Println(err)
	}
}

func (s *Service) GetItem(w http.ResponseWriter, r *http.Request) {
	itemName := chi.URLParam(r, "name")
	if itemName == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	s.RLock()
	defer s.RUnlock()
	if !s.itemExists(itemName) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(s.items[itemName])
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Service) itemExists(itemName string) bool {
	if _, ok := s.items[itemName]; ok {
		return true
	}
	return false
}
