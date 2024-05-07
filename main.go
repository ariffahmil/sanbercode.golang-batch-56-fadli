package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	var kata1, kata2, kata3, kata4, kata5 string
	kata1 = "Bootcamp"
	kata2 = "Digital"
	kata3 = "Skill"
	kata4 = "Sanbercode"
	kata5 = "Golang"

	kalimat := kata1 + " " + kata2 + " " + kata3 + " " + kata4 + " " + kata5
	fmt.Println("soal 1:")
	fmt.Println(kalimat)

	// soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println("soal 2:")
	fmt.Println(halo)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	gabungan := kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat
	fmt.Println("soal 3:")
	fmt.Println(gabungan)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angka1, _ := strconv.Atoi(angkaPertama)
	angka2, _ := strconv.Atoi(angkaKedua)
	angka3, _ := strconv.Atoi(angkaKetiga)
	angka4, _ := strconv.Atoi(angkaKeempat)

	total := angka1 + angka2 + angka3 + angka4
	fmt.Println("soal 4:")
	fmt.Println("Jumlah angka pertama, kedua, ketiga, dan keempat adalah:", total)

	// soal 5
	kalimat = "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo", "hi", 2)
	fmt.Println("soal 5:")
	fmt.Println(kalimat, "_", angka)
}
