package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"healthybank.com/healthybank/financial_project"
)

type FinancialProjectService interface {
	GetAll() []*financial_project.FinancialProject
}

type Handler struct {
	fps FinancialProjectService
}

func New(fps FinancialProjectService) *Handler {
	return &Handler{
		fps: fps,
	}
}

func (h *Handler) Start() {
	r := mux.NewRouter()
	r.HandleFunc("/financialproject", h.createProjectHandler).Methods("POST")
	r.HandleFunc("/financialproject/{id}", getByIDProjectHandler).Methods("GET")
	r.HandleFunc("/financialproject", h.getProjectHandler).Methods("GET")
	r.HandleFunc("/financialproject/{id}", updateProjectHandler).Methods("PUT")
	r.HandleFunc("/financialproject/{id}", deleteProjectHandler).Methods("DELETE")
	http.Handle("/", r)

	http.ListenAndServe(":3000", r)
}

func (h *Handler) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create projecthandler")
}

func getByIDProjectHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getProjectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get projecthandler")
	h.fps.GetAll()
}

func updateProjectHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {

}
