package library

import (
	"fmt"
	"math"
	"strings"
)

// SOAL 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}
type PersegiPanjang struct {
	Panjang, Lebar int
}
type Tabung struct {
	JariJari, Tinggi float64
}
type Balok struct {
	Panjang, Lebar, Tinggi int
}

// ////BANGUN DATAR BANGUN RUANG//////
type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}
type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

// ////SEGITIGA//////
func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}
func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

// ////PERSEGI PANJANG//////
func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}
func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

// ////TABUNG//////
func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}
func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

// ////BALOK//////
func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}
func (b Balok) LuasPermukaan() float64 {
	return 2 * float64(b.Panjang*b.Lebar+b.Panjang*b.Tinggi+b.Lebar*b.Tinggi)
}

// SOAL 2
type Phone struct {
	Name   string
	Brand  string
	Year   int
	Colors []string
}
type HP interface {
	Muncul()
}

func (p Phone) Muncul() {
	fmt.Println("\nName :", p.Name)
	fmt.Println("Brand :", p.Brand)
	fmt.Println("Year :", p.Year)
	fmt.Println("Colors :", strings.Join(p.Colors, ", "))
}

// SOAL 3
func LuasPersegi(sisi int, tampilkan bool) interface{} {
	if sisi == 0 && tampilkan {
		return "Maaf anda belum menginput sisi dari persegi"
	} else if sisi == 0 && !tampilkan {
		return nil
	}
	luas := sisi * sisi / 4

	if tampilkan {
		return fmt.Sprintf("\nLuas persegi dengan sisi 2cm adalah %dcm", luas)
	}
	return luas
}

// SOAL 4
func CetakHasil(prefix interface{}, angka1 interface{}, angka2 interface{}) {
	prefixStr := prefix.(string)
	kumpulan1 := angka1.([]int)
	kumpulan2 := angka2.([]int)

	total := 0
	semuaAngka := append(kumpulan1, kumpulan2...)
	teks := ""

	for i, val := range semuaAngka {
		total += val
		teks += fmt.Sprintf("%d", val)
		if i < len(semuaAngka)-1 {
			teks += " + "
		}
	}
	fmt.Printf("%s%s = %d\n", prefixStr, teks, total)
}

// SOAL 5
func Tambah(a, b int) int {
	return a + b
}

// Kali mengembalikan hasil perkalian dua angka
func Kali(a, b int) int {
	return a * b
}

// SOAL LIVE SESSION
func Penjumlahan(nums []int, target int) []int {
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
