package handlers

import "net/http"

type ServicesHandler struct {
}

func NewServicesHandler() *ServicesHandler {
	return &ServicesHandler{}
}

func (s *ServicesHandler) ListServices(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
}

func (s *ServicesHandler) CreateService(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(201)
}

func (s *ServicesHandler) GetService(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
}

func (s *ServicesHandler) DeleteService(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(204)
}
