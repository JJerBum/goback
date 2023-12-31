package router

import (
	"github.com/gorilla/mux"
	"goback/entgo/controller"
	"net/http"
)

func registerTodoRouoter(r *mux.Router) {
	todoRouter := r.PathPrefix("/todo").Subrouter()
	todoRouter.HandleFunc("/{id}", controller.TodoGetByIDController).Methods(http.MethodGet)
	todoRouter.HandleFunc("/", controller.TodoCreateController).Methods(http.MethodPost)
	todoRouter.HandleFunc("/{id}", controller.TodoUpdateController).Methods(http.MethodPut)
	todoRouter.HandleFunc("/{id}", controller.TodoDeleteController).Methods(http.MethodDelete)
}
