package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("GET http://127.0.0.1:8080/")
	w.Write([]byte("Hello world"))

}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("GET http://127.0.0.1:8080/login")
	w.Write([]byte("Welcome"))

}

func logout(w http.ResponseWriter, r *http.Request) {
	log.Println("GET http://127.0.0.1:8080/logout")
	w.Write([]byte("Goodbay"))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	log.Println("Run server http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
