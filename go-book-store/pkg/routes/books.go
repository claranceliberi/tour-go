package routes

import (
	"github.com/gorilla/mux"
	"github.com/claranceliberi/go-book-store/pkg/controllers"
)


var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("books/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("books/{bookId}", controllers.UpdateBook).Methods("PUT")
}