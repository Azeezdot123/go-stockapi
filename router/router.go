package router

import (
	"github.com/gorilla/mux"
	"github.com/go-stockapi/middleware"
)
func Router() *mux.Router{
	router := mux.NewRouter()
	
	router.Handlefunc("/api/stock", middleware.GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.Handlefunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.Handlefunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
}