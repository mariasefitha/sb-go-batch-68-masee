package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// FUNCTION
// SOAL 1
func kalimatt(kalimat string, tahun int) {
	fmt.Println("\n", kalimat, ",", tahun)
}

// SOAL 2 maaf mas gak ngerti gimana ngerjainnya

// SOAL 3
func tambahAngka(nilai int, angka *int) {
	*angka += nilai
}
func cetakAngka(angka *int) {
	fmt.Println("\nTotal angka:", *angka)
}

// SOAL 4
func listPhone(phones *[]string) {
	*phones = append(*phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
}

func main() {
	//SOAL 1
	defer kalimatt("Golang Backend Development", 2021)
	fmt.Println("Program dimulai...")
	fmt.Println("Melakukan proses lainnya...")
	fmt.Printf("\n")

	// SOAL 2 maaf mas gak ngerti gimana ngerjainnya

	//SOAL 3
	angka := 1
	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	//SOAL 4
	phones := []string{}
	listPhone(&phones)   //pointer
	sort.Strings(phones) //urutan
	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	} //tampilan per detik
	fmt.Printf("\n")

	//SOAL 5
	phoness := []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phoness) //bikin urutan sesuai abjad karena di outputnya diurutkan sesuai abjad
	var wg sync.WaitGroup //jalankan wait group
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("%d. %s\n", i+1, phoness[i])
		}(i)
	}
	wg.Wait()
}
