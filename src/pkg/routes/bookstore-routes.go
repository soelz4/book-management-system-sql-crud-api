package routes

import (
	"github.com/gorilla/mux"

	"go-bookstore/src/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookID}/", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookID}/", controllers.UpdateBookByID).Methods("PUT")
	router.HandleFunc("/book/{bookID}/", controllers.DeleteBookByID).Methods("DELETE")
}
