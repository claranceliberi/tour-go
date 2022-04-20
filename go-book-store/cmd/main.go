package main

import (
	"log"
	"net/http"
	"github.com/claranceliberi/go-book-store/pkg/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}