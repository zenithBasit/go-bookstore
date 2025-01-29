package controllers

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
	}
	book, _ := models.GetBookById(id)
	if book.ID == 0 {
		http.Error(w, "Book Not Found", http.StatusNotFound)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utils.ParseBody(r, &book)
	createdBook := book.CreateBook()
	json.NewEncoder(w).Encode(createdBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	deletedBook := models.DeleteBook(id)
	json.NewEncoder(w).Encode(deletedBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invaild book ID", http.StatusBadRequest)
	}
	book, _ := models.GetBookById(id)
	if book.ID == 0 {
		http.Error(w, "Book Not Found", http.StatusNotFound)
	}
	var tempBook models.Book
	utils.ParseBody(r, &tempBook)

	if tempBook.Name != "" {
		book.Name = tempBook.Name
	}
	if tempBook.Author != "" {
		book.Author = tempBook.Author
	}
	if tempBook.Publication != "" {
		book.Publication = tempBook.Publication
	}
	config.GetDB().Save(book)
	json.NewEncoder(w).Encode(book)
}
