package main

import (
	"encoding/json"
	"net/http"
)

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	products := make([]Product, 0, len(h.store.m))
	for _, v := range h.store.m {
		products = append(products, v)
	}
	h.store.RUnlock()
	jsonBytes, err := json.Marshal(products)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getProductRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	u, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		internalServerError(w, r)
		return
	}
	h.store.Lock()
	h.store.m[p.ID] = p
	h.store.Unlock()
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	matches := getProductRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		internalServerError(w, r)
		return
	}
	h.store.RLock()
	p, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	} else {
		h.store.Lock()
		h.store.m[matches[1]] = p
		h.store.Unlock()
		jsonBytes, err := json.Marshal(p)
		if err != nil {
			internalServerError(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	matches := getProductRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	_, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("product not found"))
		return
	} else {
		delete(h.store.m, matches[1])
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("product successfully deleted"))
	}

}
