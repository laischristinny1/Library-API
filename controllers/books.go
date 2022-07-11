package controllers

import (
	"Library-API/database"
	"Library-API/model"
	"Library-API/repository"
	"Library-API/requests"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Conect()
	if erro != nil {
		requests.Err(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryBooks(db)
	books, err := repository.GetAllBooks()
	if err != nil {
		requests.Err(w, http.StatusInternalServerError, err)
		return
	}
	requests.JSON(w, http.StatusOK, books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		requests.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var book model.Book
	if err = json.Unmarshal(corpoRequest, &book); err != nil {
		requests.Err(w, http.StatusBadRequest, err)
		return
	}


	db, err := database.Conect()
	if err != nil {
		requests.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryBooks(db)
	book.ID, err = repository.Create(book)
	if err != nil {
		requests.Err(w, http.StatusInternalServerError, err)
		return
	}

	requests.JSON(w, http.StatusCreated, book)

}

func SearchBook(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	bookID, erro := strconv.ParseUint(parameters["bookID"], 10, 64)
	if erro != nil {
		requests.Err(w, http.StatusBadRequest, erro)
		return
	}

	db, err := database.Conect()
	if erro != nil {
		requests.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryBooks(db)
	book, err := repository.GetByID(bookID)
	if err != nil {
		requests.Err(w, http.StatusInternalServerError, err)
		return
	}
	requests.JSON(w, http.StatusOK, book)
}

func CheckoutBook(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	bookID, erro := strconv.ParseUint(parameters["bookID"], 10, 64)
	if erro != nil {
		requests.Err(w, http.StatusBadRequest, erro)
		return
	}

	db, err := database.Conect()
	if erro != nil {
		requests.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryBooks(db)
	book, err := repository.CheckoutBook(bookID)
	if err != nil {
		requests.Err(w, http.StatusInternalServerError, err)
		return
	}

	requests.JSON(w, http.StatusOK, book)
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	bookID, erro := strconv.ParseUint(parameters["bookID"], 10, 64)
	if erro != nil {
		requests.Err(w, http.StatusBadRequest, erro)
		return
	}

	db, err := database.Conect()
	if erro != nil {
		requests.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryBooks(db)
	book, err := repository.ReturnBook(bookID)
	if err != nil {
		requests.Err(w, http.StatusInternalServerError, err)
		return
	}

	requests.JSON(w, http.StatusOK, book)
}


