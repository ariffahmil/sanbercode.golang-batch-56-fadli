package main

import (
	"fmt"
	"strings"
)

// Variabel untuk menyimpan data film
var dataFilm = []map[string]string{}

// Soal 1: Function untuk menghitung luas, keliling, dan volume persegi panjang dan balok
func luasPersegiPanjang(panjang, lebar float64) float64 {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar float64) float64 {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi float64) float64 {
	return panjang * lebar * tinggi
}

// Soal 2: Function untuk memperkenalkan seseorang dengan informasi gender, pekerjaan, dan usia
func introduce(nama, gender, pekerjaan, usia string) string {
	var prefix string
	if gender == "laki-laki" {
		prefix = "Pak "
	} else {
		prefix = "Bu "
	}
	return fmt.Sprintf("%s%s adalah seorang %s yang berusia %s tahun", prefix, nama, pekerjaan, usia)
}

// Soal 3: Function untuk menampilkan nama dan slice buah favorit seseorang
func buahFavorit(nama string, buah ...string) string {
	return fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah \"%s\"", nama, strings.Join(buah, "\", \""))
}

// Soal 4: Closure function untuk menambahkan data film ke slice dataFilm
var tambahDataFilm = func(nama, durasi, genre, tahun string) {
	dataFilm = append(dataFilm, map[string]string{
		"Nama":   nama,
		"Durasi": durasi,
		"Genre":  genre,
		"Tahun":  tahun,
	})
}

func main() {
	// Soal 1
	panjang := 12.0
	lebar := 4.0
	tinggi := 8.0

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("Luas Persegi Panjang:", luas)
	fmt.Println("Keliling Persegi Panjang:", keliling)
	fmt.Println("Volume Balok:", volume)

	// Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// Soal 3
	buah := []string{"semangka", "jeruk", "melon", "pepaya"}
	buahFavoritJohn := buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)

	// Soal 4
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
