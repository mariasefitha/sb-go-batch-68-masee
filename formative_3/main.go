package main

func main() {

	// //SOAL 1 terdapat variabel seperti di bawah ini:
	// // kalimat := "halo halo bandung"
	// // angka := 2021
	// // olah variabel diatas yang hasil output akhrinya adalah "Hi Hi bandung" - 2021
	// kalimat := "halo halo bandung"
	// angka := 2021

	// var newText1 = strings.ReplaceAll(kalimat, "halo", "Hi")
	// fmt.Println(newText1, "-", angka)

	///////////////////////////////////////////////////////////////////////////////////////////////////////

	// 	//SOAL 2 buatlah variabel seperti di bawah ini
	// 	// var nilaiJohn = 80
	// 	// var nilaiDoe = 50
	// 	// tentukan indeks nilai dari nilaiJohn dan nilaiDoe (tampilkan dengan fmt.Println) dengan kondisi :
	// 	// nilai >= 80 indeksnya A
	// 	// nilai >= 70 dan nilai < 80 indeksnya B
	// 	// nilai >= 60 dan nilai < 70 indeksnya c
	// 	// nilai >= 50 dan nilai < 60 indeksnya D
	// 	// nilai < 50 indeksnya E
	// 	var nilaiJohn = 80
	// 	var nilaiDoe = 50
	// 	fmt.Println("Indeks nilai dari John:", indeksnya(nilaiJohn))
	// 	fmt.Println("Indeks nilai dari Doe:", indeksnya(nilaiDoe))
	// }
	// func indeksnya(nilai int) string {
	// 	switch {
	// 	case nilai >= 80:
	// 		return "A"
	// 	case nilai >= 70:
	// 		fallthrough
	// 	case nilai >= 60:
	// 		if nilai < 70 {
	// 			return "C"
	// 		}
	// 		return "B"
	// 	case nilai >= 50:
	// 		fallthrough
	// 	case nilai < 50:
	// 		if nilai >= 50 {
	// 			return "D"
	// 		}
	// 		return "E"
	// 	default:
	// 		return "You need to learn more"
	// 	}
	// }

	///////////////////////////////////////////////////////////////////////////////////////////////////////
	// 	//SOAL 3 buatlah variabel seperti di bawah ini
	// 	// var tanggal = 17;
	// 	// var bulan = 8;
	// 	// var tahun = 1945;
	// 	// ganti tanggal ,bulan, dan tahun sesuai dengan tanggal lahir anda
	// 	// dan buatlah switch case pada bulan, lalu muncul kan string nya dengan output
	// 	// seperti ini 17 Agustus 1945 (isi di sesuaikan dengan tanggal lahir masing-masing)
	// 	var tanggal = 8
	// 	var bulan = 6
	// 	var tahun = 1996
	// 	fmt.Println("Tanggal Lahir:", tanggal, bulannya(bulan), tahun)
	// }
	// func bulannya(bulan int) string {

	// 	switch {
	// 	case bulan == 1:
	// 		return "Januari"
	// 	case bulan == 2:
	// 		return "Februari"
	// 	case bulan == 3:
	// 		return "Maret"
	// 	case bulan == 4:
	// 		return "April"
	// 	case bulan == 5:
	// 		return "Mei"
	// 	case bulan == 6:
	// 		return "Juni"
	// 	case bulan == 7:
	// 		return "Juli"
	// 	case bulan == 8:
	// 		return "Agustus"
	// 	case bulan == 9:
	// 		return "September"
	// 	case bulan == 10:
	// 		return "Oktober"
	// 	case bulan == 11:
	// 		return "November"
	// 	case bulan == 12:
	// 		return "Desember"
	// 	default:
	// 		return "Tidak Valid"
	// 	}
	// }

	///////////////////////////////////////////////////////////////////////////////////////////////////////
	// //SOAL 4 Berikut adalah beberapa istilah generasi berdasarkan tahun kelahirannya:
	// // Baby boomer, kelahiran 1944 s.d 1964
	// // Generasi X, kelahiran 1965 s.d 1979
	// // Generasi Y (Millenials), kelahiran 1980 s.d 1994
	// // Generasi Z, kelahiran 1995 s.d 2015
	// // buatlah conditional dimana menghasilkan istilah diatas sesuai dengan tahun kelahiran anda
	// var tahun = 1996
	// switch {
	// case (tahun >= 1944) && (tahun <= 1964):
	// 	fmt.Println("Baby boomer")
	// 	fallthrough
	// case (tahun >= 1965) && (tahun <= 1979):
	// 	fmt.Println("Generasi X")
	// 	fallthrough
	// case (tahun >= 1980) && (tahun <= 1994):
	// 	fmt.Println("Generasi Y (Millenials)")
	// 	fallthrough
	// case (tahun >= 1995) && (tahun <= 2015):
	// 	fmt.Println("Generasi Z")
	// default:
	// 	{
	// 		fmt.Println("Belum ada istilah generasi")
	// 	}
	// }

	///////////////////////////////////////////////////////////////////////////////////////////////////////
	// //SOAL 5 Pada tugas ini kamu diminta untuk melakukan looping. Untuk membuat tantangan ini lebih menarik,
	// // kamu juga diminta untuk memenuhi syarat tertentu yaitu:
	// // SYARAT:
	// // A. Jika angka ganjil maka tampilkan Santai
	// // B. Jika angka genap maka tampilkan Berkualitas
	// // C. Jika angka yang sedang ditampilkan adalah kelipatan 3 DAN angka ganjil maka tampilkan I Love Coding.
	// for i := 1; i <= 20; i++ {
	// 	if i%3 == 0 && i%2 != 0 { //angka yang sedang ditampilkan adalah kelipatan 3 DAN angka ganjil
	// 		fmt.Println("Angka ke-", i, "I Love Coding")
	// 	} else if i%2 == 0 { //angka genap
	// 		fmt.Println("Angka ke-", i, "Berkualitas")
	// 	} else { //angka ganjil
	// 		fmt.Println("Angka ke-", i, "Santai")
	// 	}
	// }
}
