package main

import "sync"

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Datastore struct {
	m map[string]Product
	*sync.RWMutex
}
