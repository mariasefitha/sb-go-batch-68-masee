package main

import (
	//SOAL 5 mohon maaf mas, saya gak ngerti gimana caranya, udah coba cari-cari tp gak ngerti dan gak bisa

	"fmt"
	"formative-9/library"
)

func main() {
	//SOAL 1
	// Bangun Datar
	segitiga := library.SegitigaSamaSisi{Alas: 4, Tinggi: 6}
	persegi := library.PersegiPanjang{Panjang: 5, Lebar: 3}

	var datar library.HitungBangunDatar

	datar = segitiga
	fmt.Println("====Segitiga Sama Sisi===")
	fmt.Println("Luas:", datar.Luas())
	fmt.Println("Keliling:", datar.Keliling())

	datar = persegi
	fmt.Println("===Persegi Panjang===")
	fmt.Println("Luas:", datar.Luas())
	fmt.Println("Keliling:", datar.Keliling())

	// Bangun Ruang
	tabung := library.Tabung{JariJari: 7, Tinggi: 10}
	balok := library.Balok{Panjang: 4, Lebar: 3, Tinggi: 5}

	var ruang library.HitungBangunRuang

	ruang = tabung
	fmt.Println("===Tabung===")
	fmt.Printf("Volume: %.2f\n", ruang.Volume())
	fmt.Printf("Luas Permukaan: %.2f\n", ruang.LuasPermukaan())

	ruang = balok
	fmt.Println("===Balok===")
	fmt.Printf("Volume: %.2f\n", ruang.Volume())
	fmt.Printf("Luas Permukaan: %.2f\n", ruang.LuasPermukaan())

	//SOAL 2
	p := library.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	var d library.HP = p
	d.Muncul()

	//SOAL 3
	fmt.Println(library.LuasPersegi(4, true))
	fmt.Println(library.LuasPersegi(8, false))
	fmt.Println(library.LuasPersegi(0, true))
	fmt.Println(library.LuasPersegi(0, false))

	//SOAL 4
	prefix := interface{}("\nHasil Penjumlahan dari: ")
	kumpulanAngkaPertama := interface{}([]int{6, 8})
	kumpulanAngkaKedua := interface{}([]int{12, 14})

	library.CetakHasil(prefix, kumpulanAngkaPertama, kumpulanAngkaKedua)

	//SOAL 5
	a := 4
	b := 5
	hasilPenjumlahan := library.Tambah(a, b)
	hasilPerkalian := library.Kali(a, b)
	fmt.Printf("Hasil Penjumlahan %d + %d = %d\n", a, b, hasilPenjumlahan)
	fmt.Printf("Hasil Perkalian %d * %d = %d", a, b, hasilPerkalian)

	//SOAL LIVE SESSION
	nums := []int{2, 7, 11, 15}
	target := 22
	result := library.Penjumlahan(nums, target)
	fmt.Println("\n")
	fmt.Println(result)
}
