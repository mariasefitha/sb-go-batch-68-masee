package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// SOAL 1 tampilkan kalimat "Bootcamp Digital Skill Sanbercode Golang" yang tersusun dari gabungan variabel dalam setiap kata (5 variabel)
	// Deklarasi variabel untuk setiap kata
	var kata1 = "Bootcamp"
	var kata2 = "Digital"
	var kata3 = "Skill"
	var kata4 = "Sanbercode"
	var kata5 = "Golang"
	// Gabungkan dan tampilkan kalimat
	var tampilankalimat = kata1 + " " + kata2 + " " + kata3 + " " + kata4 + " " + kata5
	fmt.Println(tampilankalimat)

	/////////////////////////////////////////////////////////////////////////////////////////

	// //SOAL 2 terdapat variabel seperti di bawah ini:
	// // halo := "Halo Dunia"
	// // ganti kata "Dunia" menjadi "Golang" menggunakan packages strings
	// var text = "Halo Dunia"
	// var find = "Dunia"
	// var replaceWith = "Golang"

	// var newText1 = strings.Replace(text, find, replaceWith, 1)
	// fmt.Println(newText1)
	/////////////////////////////////////////////////////////////////////////////////////////

	// SOAL 3 buatlah variabel-variabel seperti di bawah ini
	// var kataPertama = "saya";
	// var kataKedua = "senang";
	// var kataKetiga = "belajar";
	// var kataKeempat = "golang";
	// gabungkan variabel-variabel tersebut agar menghasilkan output saya Senang belajaR GOLANG
	var kataPertama = "saya"
	var kataKedua = "senang"
	var text = "belajar"
	//var kataKeempat = "golang";

	var replaceWith = "R"
	var find = "r"
	var gantikataKetiga = strings.Replace(text, find, replaceWith, 6)
	var gantikataKeempat = strings.ToUpper("golang!")

	var tampilankalimat1 = kataPertama + " " + kataKedua + " " + gantikataKetiga + " " + gantikataKeempat
	fmt.Println(tampilankalimat1)

	/////////////////////////////////////////////////////////////////////////////////////////

	// //SOAL 4 buatlahlah kode untuk variabel-variabel seperti di bawah ini
	// // var angkaPertama= "8";
	// // var angkaKedua= "5";
	// // var angkaKetiga= "6";
	// // var angkaKeempat = "7";
	// // ubahlah variabel diatas menjadi integer dan jumlahkan semuanya

	// // Variabel string angka
	// var angkaPertama = "8"
	// var angkaKedua = "5"
	// var angkaKetiga = "6"
	// var angkaKeempat = "7"

	// // Konversi string ke integer
	// var angka1, err1 = strconv.Atoi(angkaPertama)
	// var angka2, err2 = strconv.Atoi(angkaKedua)
	// var angka3, err3 = strconv.Atoi(angkaKetiga)
	// var angka4, err4 = strconv.Atoi(angkaKeempat)

	// //cek error pada setiap konversi
	// if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
	// 	fmt.Println("Terjadi kesalahan saat konversi string ke integer.")
	// 	return
	// }

	// //jumlah semua angka
	// var total = angka1 + angka2 + angka3 + angka4

	// //hasil
	// fmt.Println("Hasil Penjumlahan:", total)

	/////////////////////////////////////////////////////////////////////////////////////////

	//SOAL 5 buatlah variabel-variabel seperti di bawah ini
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"
	//ubah lah variabel diatas ke dalam integer
	var ppp, err1 = strconv.Atoi(panjangPersegiPanjang)
	var lpp, err2 = strconv.Atoi(lebarPersegiPanjang)
	var as, err3 = strconv.Atoi(alasSegitiga)
	var ts, err4 = strconv.Atoi(tinggiSegitiga)

	//cek error pada setiap konversi
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Terjadi kesalahan saat konversi string ke integer.")
		return
	}

	// dan gunakan pada operasi perhitungan dari luas persegi panjang, keliling persegi panjang dan luas segitiga dengan variabel di bawah ini:
	//var luasPersegiPanjang int
	var luasPersegiPanjang = ppp * lpp
	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)

	//var kelilingPersegiPanjang int
	var kelilingPersegiPanjang = 2 * (ppp + lpp)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)

	//var luasSegitiga int
	////var str = "0.5"
	////var num, _ := strconv.ParseFloat(str, 32)
	////var num, _ = strconv.ParseFloat(str, 32)
	var luasSegitiga = (as * ts) / 2
	//fmt.Println(num)
	fmt.Println("Luas Segitiga:", luasSegitiga)

}
