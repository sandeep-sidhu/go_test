package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var DB = NewMySQLDB()

type Book struct {
	Title  string
	Author string
	Credit string
	Year   int
}

func getRoot(resp http.ResponseWriter, r *http.Request) {
	host := r.Host
	fmt.Println("getRoot - request received on / from host: ", host)
	resp.Write([]byte("root"))
}

func getHello(blah http.ResponseWriter, r *http.Request) {
	host := r.Host
	headers := r.Header.Get("client-id")
	fmt.Println("getHello - request received /hello from host: ", host)
	fmt.Println("client-id: ", headers)
	blah.Write([]byte("Hello"))
}

func getBook(resp http.ResponseWriter, r *http.Request) {

	book, err := DB.GetBook(1)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	// Convert the book struct to a JSON string
	bookJson, err := json.Marshal(book)
	// Print bookJson
	fmt.Println(string(bookJson))

	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("client-id", "blah blah blah")
	resp.Write(bookJson)
}

func main() {
	// Create a new MySQLDB instance
	DB.Connect("username:password@tcp(127.0.0.1:3306)/dbname")

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/book", getBook)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
