package main

import (
	//SOAL 5 mohon maaf mas, saya gak ngerti gimana caranya, udah coba cari-cari tp gak ngerti dan gak bisa

	"fmt"
	"math"
)

// Deklarasi Struct dan Interface//
// SOAL 1
type segitigaSamaSisi struct {
	sisi int
}
type persegiPanjang struct {
	panjang, lebar int
}
type tabung struct {
	jariJari, tinggi float64
}
type balok struct {
	panjang, lebar, tinggi int
}

// SOAL 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func main() {
	//SOAL 1
	// Bangun Datar (Segitiga, Persegi Panjang)
	segitiga := segitigaSamaSisi{sisi: 6}
	persegi := persegiPanjang{panjang: 4, lebar: 2}

	fmt.Println("===Bangun Datar===")
	fmt.Println("-Luas Segitiga:", segitiga.luas())
	fmt.Println("-Keliling Segitiga:", segitiga.keliling())
	fmt.Println("-Luas Persegi Panjang:", persegi.luas())
	fmt.Println("-Keliling Persegi Panjang:", persegi.keliling())

	// Bangun Ruang (Tabung, Balok)
	tabung1 := tabung{jariJari: 5, tinggi: 7}
	balok1 := balok{panjang: 4, lebar: 5, tinggi: 6}

	fmt.Println("\n===Bangun Ruang===")
	fmt.Printf("Volume Tabung: %.2f\n", tabung1.volume())
	fmt.Printf("-Luas Permukaan Tabung: %.2f\n", tabung1.luasPermukaan())
	fmt.Printf("-Volume Balok: %.2f\n", balok1.volume())
	fmt.Printf("-Luas Permukaan Balok: %.2f\n", balok1.luasPermukaan())

	//SOAL 2
	hp := phone{
		name:  "Samsung Galaxy Note 20",
		brand: "Samsung Galaxy Note 20",
		year:  2020,
	}
	hp.warnaHP("Mystic Bronz, Mystic White, Mystic Black")
	fmt.Println("\nNama:", hp.name)
	fmt.Println("Brand:", hp.brand)
	fmt.Println("Tahun:", hp.year)
	fmt.Println("Warna:", hp.colors)

	//SOAL 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	//SOAL 4
	var prefix interface{} = "\nHasil penjumlahan dari " //interface kosong
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	angka1 := kumpulanAngkaPertama.([]int) //type assertion nya disini
	angka2 := kumpulanAngkaKedua.([]int)   //type assertion nya disini

	total := 0
	output := prefix.(string)

	for i, val := range angka1 {
		total += val
		output += fmt.Sprintf("%d", val)
		if i != len(angka1)-1 || len(angka2) > 0 {
			output += " + "
		}
	}
	for i, val := range angka2 {
		total += val
		output += fmt.Sprintf("%d", val)
		if i != len(angka2)-1 {
			output += " + "
		}
	}
	output += fmt.Sprintf(" = %d", total)
	fmt.Println(output)

	//SOAL 5 mohon maaf mas, saya gak ngerti gimana caranya, udah coba cari-cari tp gak ngerti dan gak bisa
	// a := 5
	// b := 3
	// hasilTambah := matematika.Tambah(a, b)
	// hasilKali := matematika.Kali(a, b)
	// fmt.Println("Hasil Tambah:", hasilTambah) // Output: 8
	// fmt.Println("Hasil Kali:", hasilKali)     // Output: 15

	//SOAL LIVE SESSION
	nums := []int{2, 7, 11, 15}
	target := 22

	result := penjumlahan(nums, target)
	fmt.Println("\n")
	fmt.Println(result)
}

// FUNCTION
// SOAL 1
func (s segitigaSamaSisi) luas() int {
	return s.sisi * s.sisi / 2
}
func (s segitigaSamaSisi) keliling() int {
	return s.sisi * 3
}
func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}
func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}
func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}
func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}
func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}
func (b balok) luasPermukaan() float64 {
	p, l, t := float64(b.panjang), float64(b.lebar), float64(b.tinggi)
	return 2 * (p*l + p*t + l*t)
}

// SOAL 2
func (p *phone) warnaHP(warna string) {
	p.colors = append(p.colors, warna)
}

// SOAL 3
func luasPersegi(sisi int, hasil bool) interface{} { //interface kosong
	if sisi == 0 && hasil {
		return "Maaf anda belum menginput sisi dari persegi"
	} else if sisi == 0 && !hasil {
		return "nil"
	} else if hasil {
		luas := sisi * sisi
		return fmt.Sprintf("\nLuas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	} else {
		return sisi * sisi
	}
}

// SOAL LIVE SESSION
func penjumlahan(nums []int, target int) []int {
	//your code here
	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
