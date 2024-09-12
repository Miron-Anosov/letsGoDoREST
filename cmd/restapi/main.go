package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	// Полуваем все заметки
	log.Println("GET status 200 http://127.0.0.1:8080/login")
	w.Write([]byte("Welcome"))

}

func get_notes_by_id(w http.ResponseWriter, r *http.Request) {
	// Получаем заметку по ID
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		log.Printf("GET status 400 http://127.0.0.1:8080/ params: %v.", id)
		http.NotFound(w, r)
		return
	}

	log.Printf("GET status 200 http://127.0.0.1:8080/notes?id=%d", id)
	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)

}

func post_new_note(w http.ResponseWriter, r *http.Request) {
	// Создает новую заметку
	if r.Method != http.MethodPost {
		log.Println("POST status 405 http://127.0.0.1:8080/notes")
		w.Header().Set("Allow", http.MethodPost) // устанавливаем заголовок ответа.
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json") // сначало нужно установить заголовокБ а потом статус-код
	w.WriteHeader(201)                                 //  вызывать можно только один раз перед отправкой ответа.
	w.Write([]byte(`{"status": OK}`))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/notes", get_notes_by_id)
	mux.HandleFunc("/notes/new", post_new_note)
	log.Println("Run server http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
