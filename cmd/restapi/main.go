package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	url := &r.URL.Path
	if *url != "/" {
		log.Printf("GET status 404 http://127.0.0.1:8080%s", *url)
		http.NotFound(w, r)
		return
	}
	log.Println("GET status 200 http://127.0.0.1:8080/")
	w.Write([]byte("Hello world"))

}

func login(w http.ResponseWriter, r *http.Request) {
	// вход
	log.Println("GET status 200 http://127.0.0.1:8080/login")
	w.Write([]byte("Welcome"))

}

func logout(w http.ResponseWriter, r *http.Request) {
	// выход из учетки
	log.Println("GET status 200 http://127.0.0.1:8080/logout")
	w.Write([]byte("Goodbay"))

}

func notes(w http.ResponseWriter, r *http.Request) {
	// Создает новую заметку
	if r.Method != http.MethodPost {
		log.Println("POST status 405 http://127.0.0.1:8080/notes")
		w.WriteHeader(405) //  вызывать можно только один раз перед отправкой ответа.
		w.Write([]byte("You can use only POST method"))
		return
	}

	w.WriteHeader(201)
	w.Write([]byte("Create a new notes"))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/notes", notes)
	log.Println("Run server http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
