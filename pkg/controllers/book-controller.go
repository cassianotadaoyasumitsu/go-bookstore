package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/cassianotadaoyasumitsu/go-bookstore/pkg/models"
	"github.com/cassianotadaoyasumitsu/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	write, err := w.Write(res)
	if err != nil {
		fmt.Println("Error while writing response GetBook")
	}
	write = write
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing GetBookById")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	write, err := w.Write(res)
	if err != nil {
		fmt.Println("Error while writing response GetBookById")
	}
	write = write
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	write, err := w.Write(res)
	if err != nil {
		fmt.Println("Error while writing response CreateBook")
	}
	write = write
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing DeleteBook")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	write, err := w.Write(res)
	if err != nil {
		fmt.Println("Error while writing response DeleteBook")
	}
	write = write
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing UpdateBook")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	write, err := w.Write(res)
	if err != nil {
		fmt.Println("Error while writing response UpdateBook")
	}
	write = write
}
