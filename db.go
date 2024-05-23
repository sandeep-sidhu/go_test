package main

import "database/sql"

type MySQLDB struct {
	db *sql.DB
}

func NewMySQLDB() *MySQLDB {
	return &MySQLDB{}
}

func (m *MySQLDB) Connect(dsn string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	m.db = db
	return nil
}

func (m *MySQLDB) Close() error {
	return m.db.Close()
}

func (m *MySQLDB) GetBook(bookID int) (Book, error) {
	// Query the database for a book with the given ID
	// and return it as a Book struct
	query := "SELECT title, author, credit, year FROM books WHERE id = ?"
	row := m.db.QueryRow(query, bookID)
	var book Book
	err := row.Scan(&book.Title, &book.Author, &book.Credit, &book.Year)
	if err != nil {
		return Book{}, err
	}
	return Book{}, nil
}
