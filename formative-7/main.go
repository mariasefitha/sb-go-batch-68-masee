package main

import "fmt"

//Deklarasi Struct//
//SOAL 1
type mahasiswa struct {
	nama, NIM string
	usia      int
}

//SOAL 2
type segitiga struct {
	alas, tinggi int
}
type persegi struct {
	sisi int
}
type persegiPanjang struct {
	panjang, lebar int
}

//SOAL 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

//SOAL 4
type film struct {
	judul, genre  string
	durasi, tahun int
}

//SOAL 5
type Hewan interface {
	Suara() string
}
type Kucing struct{}
type Anjing struct{}

func main() {
	//SOAL 1
	var mase mahasiswa
	mase.nama = "Maria Sefitha"
	mase.NIM = "v00027335"
	mase.usia = 29

	fmt.Println("Nama Mahasiswa  :", mase.nama)
	fmt.Println("NIM :", mase.NIM)
	fmt.Println("Usia :", mase.usia)

	//SOAL 2
	s := segitiga{alas: 10, tinggi: 5}
	p := persegi{sisi: 2}
	pp := persegiPanjang{panjang: 5, lebar: 3}

	fmt.Printf("\nLuas Segitiga: %.2f\n", s.luas())
	fmt.Println("Luas Persegi:", p.luas())
	fmt.Println("Luas Persegi Panjang:", pp.luas())

	//SOAL 3
	// Membuat objek phone
	hp := phone{
		name:  "Galaxy S24FE",
		brand: "Samsung",
		year:  2024,
	}

	hp.warnaHP("Black")
	hp.warnaHP("Silver")
	hp.warnaHP("Blue")

	fmt.Println("\nNama:", hp.name)
	fmt.Println("Brand:", hp.brand)
	fmt.Println("Tahun:", hp.year)
	fmt.Println("Warna:", hp.colors)
	fmt.Println("\n")

	//SOAL 4
	var dataFilm = []film{}

	tambahDataFilm("LOTR", "action", 120, 1999, &dataFilm)
	tambahDataFilm("avenger", "action", 120, 2019, &dataFilm)
	tambahDataFilm("spiderman", "action", 120, 2004, &dataFilm)
	tambahDataFilm("juon", "horror", 120, 2004, &dataFilm)

	for i, tampilFilm := range dataFilm {
		fmt.Printf("%d. title: %s\n", i+1, tampilFilm.judul)
		fmt.Printf("   duration: %d menit \n", tampilFilm.durasi)
		fmt.Printf("   genre: %s\n", tampilFilm.genre)
		fmt.Printf("   year: %d\n", tampilFilm.tahun)
	}

	//SOAL 5
	var h Hewan
	h = Kucing{}
	fmt.Println("\n")
	fmt.Println("Kucing suaranya:", h.Suara())
	h = Anjing{}
	fmt.Println("Anjing suaranya:", h.Suara())

}

//FUNCTION
//SOAL 2
func (s segitiga) luas() float64 {
	return (float64(s.alas) * float64(s.tinggi)) / 2
}
func (p persegi) luas() int {
	return p.sisi * p.sisi
}
func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

//SOAL 3
func (p *phone) warnaHP(warna string) {
	p.colors = append(p.colors, warna)
}

//SOAL 4 menggunakan pointer
func tambahDataFilm(judul, genre string, durasi, tahun int, dataFilm *[]film) {
	tampilFilm := film{
		judul:  judul,
		genre:  genre,
		durasi: durasi,
		tahun:  tahun,
	}
	*dataFilm = append(*dataFilm, tampilFilm)
}

//SOAL 5
func (k Kucing) Suara() string {
	return "Meowww~~"
}
func (a Anjing) Suara() string {
	return "Guk guk guk"
}
