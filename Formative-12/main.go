package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// struktur data untuk Nilai Mahasiswa
type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
	IndeksNilai string `json:"indeks_nilai"`
}

var dataNilai = []NilaiMahasiswa{}
var idCounter uint = 1

// MIDDLEWARE
func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ini dari middleware Log....\n")
		username, password, ok := r.BasicAuth()
		if !ok || username != "admin" || password != "mase1234" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Username atau Password tidak sesuai"))
			return
		}
		next(w, r)
	}
}

// buatlah API POST Nilai Mahasiswa untuk menambahkan data NilaiMahasiswa
// POST
func tambahNilaiArray(w http.ResponseWriter, r *http.Request) {
	var inputs []NilaiMahasiswa

	err := json.NewDecoder(r.Body).Decode(&inputs)
	if err != nil {
		http.Error(w, "Pengisian data tidak valid!", http.StatusBadRequest)
		return
	}

	for i := range inputs {
		if inputs[i].Nilai > 100 {
			http.Error(w, fmt.Sprintf("Nilai tidak boleh lebih dari 100: %s", inputs[i].Nama), http.StatusBadRequest)
			return
		}

		//Pehitungan nilai diubah ke indeks
		switch {
		case inputs[i].Nilai >= 80:
			inputs[i].IndeksNilai = "A"
		case inputs[i].Nilai >= 70:
			inputs[i].IndeksNilai = "B"
		case inputs[i].Nilai >= 60:
			inputs[i].IndeksNilai = "C"
		case inputs[i].Nilai >= 50:
			inputs[i].IndeksNilai = "D"
		default:
			inputs[i].IndeksNilai = "E"
		}
		inputs[i].ID = idCounter
		idCounter++ //otomatis terisi apabila data yang diinput lebih dari 1
		dataNilai = append(dataNilai, inputs[i])
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inputs)
}

// buatlah API GET Nilai Mahasiswa untuk menampilkan semua data NilaiMahasiswa
// GET
func tampilNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataNilai)
}

// Main function nya
func main() {
	http.HandleFunc("/nilai/mahasiswa", basicAuth(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			tambahNilaiArray(w, r)
		case "GET":
			tampilNilai(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	fmt.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
