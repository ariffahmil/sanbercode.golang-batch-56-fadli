package main

import "fmt"

func main() {
	// Soal 1
	var panjangPersegiPanjang = 8
	var lebarPersegiPanjang = 5

	var alasSegitiga = 6
	var tinggiSegitiga = 7

	// Perhitungan luas dan keliling persegi panjang
	luasPersegiPanjang := panjangPersegiPanjang * lebarPersegiPanjang
	kelilingPersegiPanjang := 2 * (panjangPersegiPanjang + lebarPersegiPanjang)

	// Perhitungan luas segitiga
	luasSegitiga := (alasSegitiga * tinggiSegitiga) / 2

	// Menampilkan hasil perhitungan
	fmt.Println("Luas persegi panjang:", luasPersegiPanjang)
	fmt.Println("Keliling persegi panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas segitiga:", luasSegitiga)

	// Soal 2
	nilaiJohn := 80
	nilaiDoe := 50

	// Menentukan indeks nilai
	fmt.Print("Indeks nilai John: ")
	if nilaiJohn >= 80 {
		fmt.Println("A")
	} else if nilaiJohn >= 70 {
		fmt.Println("B")
	} else if nilaiJohn >= 60 {
		fmt.Println("C")
	} else if nilaiJohn >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	fmt.Print("Indeks nilai Doe: ")
	if nilaiDoe >= 80 {
		fmt.Println("A")
	} else if nilaiDoe >= 70 {
		fmt.Println("B")
	} else if nilaiDoe >= 60 {
		fmt.Println("C")
	} else if nilaiDoe >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	// Soal 3
	tanggal := 10
	bulan := 12
	tahun := 2000

	// Switch case untuk bulan
	fmt.Print("Tanggal lahir: ")
	fmt.Print(tanggal)
	switch bulan {
	case 1:
		fmt.Print(" Januari ")
	case 2:
		fmt.Print(" Februari ")
	case 3:
		fmt.Print(" Maret ")
	case 4:
		fmt.Print(" April ")
	case 5:
		fmt.Print(" Mei ")
	case 6:
		fmt.Print(" Juni ")
	case 7:
		fmt.Print(" Juli ")
	case 8:
		fmt.Print(" Agustus ")
	case 9:
		fmt.Print(" September ")
	case 10:
		fmt.Print(" Oktober ")
	case 11:
		fmt.Print(" November ")
	case 12:
		fmt.Print(" Desember ")
	}
	fmt.Println(tahun)

	// Soal 4
	tahunLahir := 2000

	// Menentukan generasi berdasarkan tahun kelahiran
	fmt.Print("Anda termasuk dalam generasi: ")
	if tahunLahir >= 1944 && tahunLahir <= 1964 {
		fmt.Println("Baby boomer")
	} else if tahunLahir >= 1965 && tahunLahir <= 1979 {
		fmt.Println("Generasi X")
	} else if tahunLahir >= 1980 && tahunLahir <= 1994 {
		fmt.Println("Generasi Y (Millennials)")
	} else if tahunLahir >= 1995 && tahunLahir <= 2015 {
		fmt.Println("Generasi Z")
	} else {
		fmt.Println("Tidak termasuk dalam kategori yang diberikan")
	}
}
