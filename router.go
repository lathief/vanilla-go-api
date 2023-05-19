package main

import (
	"net/http"
	"regexp"
)

var (
	listProductRe   = regexp.MustCompile(`^\/products[\/]*$`)
	getProductRe    = regexp.MustCompile(`^\/products\/(\d+)$`)
	createProductRe = regexp.MustCompile(`^\/products[\/]*$`)
	updateProductRe = regexp.MustCompile(`^\/products\/(\d+)$`)
	deleteProductRe = regexp.MustCompile(`^\/products\/(\d+)$`)
)

func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listProductRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getProductRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createProductRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	case r.Method == http.MethodPut && updateProductRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	case r.Method == http.MethodDelete && deleteProductRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}
