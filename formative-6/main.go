package main

import (
	"fmt"
)

func main() {
	//SOAL 1 Tulislah sebuah function dengan nama introduce
	var sentence string
	introduce(&sentence, pengganti("John"), pengganti("laki-laki"), pengganti("penulis"), pengganti("30"))
	fmt.Println(sentence) // Pak John adalah seorang penulis yang berusia 30 tahun

	introduce(&sentence, pengganti("Sarah"), pengganti("perempuan"), pengganti("model"), pengganti("28"))
	fmt.Println(sentence) // Bu Sarah adalah seorang model yang berusia 28 tahun

	//SOAL 2
	var buah = []string{}
	tambahBuah(&buah)

	for i, b := range buah {
		fmt.Printf("%d. %s\n", i+1, b)
	}

	//SOAL 3
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, tampilFilm := range dataFilm {
		fmt.Printf("\n%d. title: %s\n", i+1, tampilFilm["title"])
		fmt.Printf("   duration: %s\n", tampilFilm["duration"])
		fmt.Printf("   genre: %s\n", tampilFilm["genre"])
		fmt.Printf("   year: %s\n", tampilFilm["year"])
	}

	//SOAL 4
	masukkanAngka := [10]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("\nAngka sebelum diubah (ganjil):", masukkanAngka)

	ubahJadiGenap(&masukkanAngka)
	fmt.Println("Angka setelah diubah (genap):", masukkanAngka)

	//SOAL 5 mohon maaf mas, saya gak ngerti gimana caranya
	// Membuat data buah sesuai tabel
	//Buah := 0
	// 	for tableBuah := bool daftarBuah{
	// 	nanas := Buah{"Nanas", "Kuning", false, 9000}
	// 	jeruk := Buah{"Jeruk", "Oranye", true, 8000}
	// 	semangka := Buah{"Semangka", "Hijau & Merah", false, 10000}
	// 	pisang := Buah{"Pisang", "Kuning", true, 5000}
}

// SOAL 1
func introduce(sentence *string, nama *string, jenisKelamin *string, pekerjaan *string, umur *string) {
	switch *jenisKelamin {
	case "laki-laki":
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", *nama, *pekerjaan, *umur)
	case "perempuan":
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun \n", *nama, *pekerjaan, *umur)
	default:
		*sentence = fmt.Sprintf("Pak/Bu %s adalah seorang %ps yang berusia %s tahun", *nama, *pekerjaan, *umur)
	}
}
func pengganti(s string) *string {
	return &s
}

// SOAL 2
func tambahBuah(buah *[]string) {
	*buah = append(*buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
}

// SOAL 3
func tambahDataFilm(judul, durasi, jenis, tahun string, dataFilm *[]map[string]string) {
	tampilFilm := map[string]string{
		"title":    judul,
		"duration": durasi,
		"genre":    jenis,
		"year":     tahun,
	}
	*dataFilm = append(*dataFilm, tampilFilm)
}

// SOAL 4
func ubahJadiGenap(arr *[10]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

// SOAL 5 mohon maaf mas, saya gak ngerti gimana caranya
// func buah(namaBuah, warnaBuah, harga string, AdaBijinya bool, daftarBuah *[]map[string]string) {
// 	buah := map[string]string{
// 		"Nama":        namaBuah,
// 		"Warna":       warnaBuah,
// 		"Ada Bijinya": AdaBijinya,
// 		"Harga":       harga,
// 	}
// 	*daftarBuah = append(*daftarBuah, buah)
// }
