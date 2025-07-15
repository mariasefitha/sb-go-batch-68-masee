package main

import "fmt"

func main() {

	//SOAL 1 Kali ini kamu diminta untuk menampilkan sebuah segitiga dengan tanda pagar (#)
	//// dengan dimensi tinggi 7 dan alas 7. gunakan looping untuk mengerjakan soal ini
	t := 7

	for i := 1; i <= t; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// //SOAL 2 buatlah variabel seperti di bawah ini
	// // var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	// // olah variabel diatas agar menghasilkan output seperti dibawah ini
	// // [saya sangat senang belajar golang]
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var newkalimat = kalimat[2:7]
	fmt.Println(newkalimat)

	// //SOAL 3 buatlah variabel seperti di bawah ini
	// //var i = 0
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, sayur := range sayuran {
		fmt.Printf("%d. %s\n", i+1, sayur)
	}

	// //SOAL 4 buatlah variabel seperti di bawah ini
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	var volumeBalok = (satuan["panjang"] * satuan["lebar"] * satuan["tinggi"])

	fmt.Println("Panjang = ", satuan["panjang"])
	fmt.Println("Lebar = ", satuan["lebar"])
	fmt.Println("Tinggi = ", satuan["tinggi"])
	fmt.Println("Volume Balok =", volumeBalok)

	// //SOAL 5 Tulislah 3 function dengan nama luas persegi panjang, keliling persegi panjang dan volume balok
	var ukuran = map[string]int{
		"panjangg": 12,
		"lebarr":   4,
		"tinggii":  4,
	}
	luasPersegiPanjang := ukuran["panjangg"] * ukuran["lebarr"]
	kelilingPersegiPanjang := 2 * (ukuran["panjangg"] + ukuran["lebarr"])
	volumeBalokk := (ukuran["panjangg"] * ukuran["lebarr"] * ukuran["tinggii"])

	fmt.Println("Luas =", luasPersegiPanjang)
	fmt.Println("Keliling =", kelilingPersegiPanjang)
	fmt.Println("Volume =", volumeBalokk)

}
