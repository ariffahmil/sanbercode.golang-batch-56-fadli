package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var (
	nilaiNilaiMahasiswa = []NilaiMahasiswa{}
	mutex               sync.Mutex
	nextID              uint = 1
)

// BasicAuth middleware function
func BasicAuth(next http.HandlerFunc, username, password string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || u != username || p != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Handler for POST /nilai
func postNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama, MataKuliah string
		Nilai            uint
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if input.Nilai > 100 {
		http.Error(w, "Nilai cannot be more than 100", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	indeksNilai := getIndeksNilai(input.Nilai)
	nilaiMahasiswa := NilaiMahasiswa{
		Nama		: input.Nama,
		MataKuliah	: input.MataKuliah,
		Nilai		: input.Nilai,
		ID			: nextID,
		IndeksNilai	: indeksNilai,
	}

	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)
	nextID++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nilaiMahasiswa)
}

// Handler for GET /nilai
func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

// Helper function to determine IndeksNilai
func getIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

func main() {
	http.HandleFunc("/nilai", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			BasicAuth(postNilaiMahasiswa, "admin", "password123")(w, r)
		case http.MethodGet:
			getNilaiMahasiswa(w, r)
		default:
			http.Error(w, "Only GET and POST methods are allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running at http://localhost:8080/nilai...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
