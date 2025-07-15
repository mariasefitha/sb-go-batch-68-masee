package main

import (
	"fmt"
	"math"
	"net/http"
)

func volumeTabung(jariJari, tinggi float64) float64 {
	return math.Pi * jariJari * jariJari * tinggi
}
func luasAlas(jariJari float64) float64 {
	return math.Pi * jariJari * jariJari
}
func kelilingAlas(jariJari float64) float64 {
	return 2 * math.Pi * jariJari
}
func index(w http.ResponseWriter, r *http.Request) {
	jariJari := 7.00
	tinggi := 10.00
	volume := volumeTabung(jariJari, tinggi)
	luasAlas := luasAlas(jariJari)
	kelilingAlas := kelilingAlas(jariJari)

	perhitungan := fmt.Sprintf(
		"Jari-jari: %.0f, Tinggi: %.0f, Volume: %.2f, Luas Alas: %.2f, Keliling Alas: %.2f",
		jariJari, tinggi, volume, luasAlas, kelilingAlas,
	)
	fmt.Fprintln(w, perhitungan)
}

func main() {
	http.HandleFunc("/test", index)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
