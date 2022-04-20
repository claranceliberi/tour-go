package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/claranceliberi/go-book-store/pkg/models"
	"github.com/claranceliberi/go-book-store/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Can't parse id of %s", bookId)
		fmt.Println(err)
	}

	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r, Book)

	book := Book.Create()

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Can't parse id of %s", bookId)
		fmt.Println(err)
	}

	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Can't parse id of %s", bookId)
		fmt.Println(err)
	}

	foundBook, db := models.GetBookById(id)


	if UpdateBook.Name != ""{
		foundBook.Name = UpdateBook.Name
	}
	if UpdateBook.Author != ""{
		foundBook.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != ""{
		foundBook.Publication = UpdateBook.Publication
	}

	db.Save(&foundBook)

	res, _ := json.Marshal(foundBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
