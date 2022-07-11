package repository

import (
	"Library-API/model"
	"database/sql"
)

type Books struct {
	db *sql.DB
}

func NewRepositoryBooks(db *sql.DB) BookRepository {
	return &Books{db}
}

func (repository Books) Create(book model.Book) (uint64, error){
	statement, err := repository.db.Prepare("insert into books (title, author, quantity) values(?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(book.Title, book.Author, book.Quantity)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

func (repository Books) GetByID(ID uint64) (model.Book, error){
	lines, err := repository.db.Query(
		"select id, title, author, quantity from books where id = ?",
		ID,
	)
	if err != nil {
		return model.Book{}, err
	}
	defer lines.Close()

	var book model.Book

	if lines.Next() {
		if err = lines.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Quantity,
		); err != nil {
			return model.Book{}, err
		}
	}

	return book, nil
}

func (repository Books) GetAllBooks () ([]model.Book, error) {
	
	linhas, erro := repository.db.Query(
		"select * from books",
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var books []model.Book

	for linhas.Next() {
		var book model.Book

		if erro = linhas.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Quantity,
		); erro != nil {
			return nil, erro
		}

		books = append(books, book)
	}

	return books, nil
}

func (repository Books) CheckoutBook(ID uint64) (model.Book, error){

	lines2, err := repository.db.Query(
		"update books set quantity = quantity-1",
	)
	if err != nil {
		return model.Book{}, err
	}
	defer lines2.Close()
	
	lines, err := repository.db.Query(
		"select id, title, author, quantity from books where id = ?",
		ID,
	)
	if err != nil {
		return model.Book{}, err
	}
	defer lines.Close()

	var book model.Book
	if lines.Next() {
		if err = lines.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Quantity,
		); err != nil {
			return model.Book{}, err
		}
	}
	
	return book, nil
}

func (repository Books) ReturnBook(ID uint64) (model.Book, error){

	lines2, err := repository.db.Query(
		"update books set quantity = quantity+1",
	)
	if err != nil {
		return model.Book{}, err
	}
	defer lines2.Close()
	
	lines, err := repository.db.Query(
		"select id, title, author, quantity from books where id = ?",
		ID,
	)
	if err != nil {
		return model.Book{}, err
	}
	defer lines.Close()

	var book model.Book
	if lines.Next() {
		if err = lines.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Quantity,
		); err != nil {
			return model.Book{}, err
		}
	}
	
	return book, nil
}
