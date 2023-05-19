package main

import (
	"fmt"
	"net/http"
	"sync"
)

type ProductHandler struct {
	store *Datastore
}

func main() {
	mux := http.NewServeMux()
	productH := &ProductHandler{
		store: &Datastore{
			m: map[string]Product{
				//initial data
				"1": Product{ID: "1", Name: "Samsung", Description: "SmartPhone", Price: 10000000},
				"2": Product{ID: "2", Name: "Nokia", Description: "SmartPhone", Price: 4000000},
			},
			RWMutex: &sync.RWMutex{},
		},
	}
	mux.Handle("/products", productH)
	mux.Handle("/products/", productH)
	fmt.Println("running on port 8080")
	http.ListenAndServe("localhost:8080", mux)
}
