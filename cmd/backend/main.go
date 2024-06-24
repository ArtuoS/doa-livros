package main

import (
	"log"
	"net/http"

	"github.com/ArtuoS/doa-livros/internal/controller"
	"github.com/ArtuoS/doa-livros/internal/repository"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=doalivros sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	bookRepo := &repository.BookRepository{DB: db}
	userRepo := &repository.UserRepository{DB: db}
	donatedBookRepo := &repository.DonatedBookRepository{DB: db}

	userController := controller.NewUserController(userRepo)
	bookController := controller.NewBookController(bookRepo, donatedBookRepo)

	r := mux.NewRouter()

	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	r.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", bookController.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/retrieve", bookController.RedeemBook).Methods("POST")

	log.Println("Starting server on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
