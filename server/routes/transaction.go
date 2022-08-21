package routes

import (
	"waysbucks/handlers"
	// "waysbucks/pkg/middleware"
	"waysbucks/pkg/mysql"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func Transaction(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", (h.FindTransactions)).Methods("GET")
	r.HandleFunc("/transaction/{id}", (h.GetTransaction)).Methods("GET")
	r.HandleFunc("/transaction", (h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction/{id}", (h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", (h.DeleteTransaction)).Methods("DELETE")
}
