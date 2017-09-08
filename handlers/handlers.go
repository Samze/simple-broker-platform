package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	})

	servicesHandler := NewServicesHandler()

	r.HandleFunc("/services", servicesHandler.CreateService).Methods("POST")
	r.HandleFunc("/services", servicesHandler.ListServices).Methods("GET")
	r.HandleFunc("/services/{guid}", servicesHandler.GetService).Methods("GET")
	r.HandleFunc("/services/{guid}", servicesHandler.DeleteService).Methods("DELETE")

	return r
}
