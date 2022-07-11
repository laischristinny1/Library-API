package repository

import "Library-API/model"

type BookRepository interface {
	Create(book model.Book) (uint64, error)
	GetByID(ID uint64) (model.Book, error)
	GetAllBooks () ([]model.Book, error) 
	CheckoutBook(ID uint64) (model.Book, error)
	ReturnBook(ID uint64) (model.Book, error)
}