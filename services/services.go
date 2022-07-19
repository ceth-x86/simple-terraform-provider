package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Entity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/entities", func(w http.ResponseWriter, r *http.Request) {

		enitity := Entity{
			Id:          "001",
			Name:        "entity_001",
			Description: "sample entity",
		}
		var entities []Entity
		entities = append(entities, enitity)

		data, err := json.Marshal(entities)
		if err != nil {
			log.Println(err)
		}

		w.Write(data)
	})
	http.ListenAndServe(":3000", r)
}
