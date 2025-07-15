package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//SOAL 1
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah, "\n") // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//SOAL 2
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn, "\n")

	//SOAL 3
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	tambahDataFilm := func(genre, jam, tahun, title string) {
		film := map[string]string{
			"genre": genre,
			"jam":   jam,
			"tahun": tahun,
			"title": title,
		}
		dataFilm = append(dataFilm, film)
	}
	tambahDataFilm("action", "2 jam", "1999", "LOTR")
	tambahDataFilm("action", "2 jam", "2019", "avenger")
	tambahDataFilm("action", "2 jam", "2004", "spiderman")
	tambahDataFilm("horror", "2 jam", "2004", "juon")

	for _, filmm := range dataFilm {
		fmt.Println(filmm, "\n")
	}

	//SOAL 4 maaf saya gak bisa ngerjain ini, gak paham soalnya:(

	//SOAL 5
	var luasLingkaran float64
	var kelilingLingkaran float64
	var jariJari float64 = 8

	perhitungan(&luasLingkaran, &kelilingLingkaran, jariJari)

	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)
}

// SOAL 1 Tulislah sebuah function dengan nama introduce yang menggunakan predefined value/ named value.
// pastikan semua parameter pada function introduce di gunakan semuanya
func introduce(nama, jenis, pekerjaan, umur string) string {
	var jenisKelamin string
	if jenis == "laki-laki" {
		jenisKelamin = "Pak"
	} else if jenis == "perempuan" {
		jenisKelamin = "Bu"
	} else {
		jenis = ""
	}
	return fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", jenisKelamin, nama, pekerjaan, umur)
}

// SOAL 2 buatlah function yang menampung data slice di bawah ini sebagai variadic function
func buahFavorit(nama string, buah ...string) string {
	hasilBuah := "\"" + strings.Join(buah, "\", \"") + "\""
	return fmt.Sprintf("Halo Nama Saya %s dan Buah Favorit Saya Adalah %s", strings.ToLower(nama), hasilBuah)
}

// SOAL 4 maaf saya gak bisa ngerjain ini, gak paham soalnya:(

// SOAL 5 Tulislah kode seperti di bawah ini, buatlah var luasLigkaran float64 var kelilingLingkaran float64
func perhitungan(rumusLuasLingkaran *float64, rumusKelilingLingkaran *float64, jariJari float64) {
	*rumusLuasLingkaran = math.Pi * jariJari * jariJari
	*rumusKelilingLingkaran = 2 * math.Pi * jariJari
}
