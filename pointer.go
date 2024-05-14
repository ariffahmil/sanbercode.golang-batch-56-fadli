package main

import (
	"fmt"
	"math"
)

// Soal 1
var luasLingkaran float64
var kelilingLingkaran float64

func hitungLingkaran(jariJari float64, luas *float64, keliling *float64) {
	*luas = math.Pi * jariJari * jariJari
	*keliling = 2 * math.Pi * jariJari
}

// Soal 2
var sentence string

func introduce(sentence *string, name string, gender string, profession string, age string) {
	if gender == "laki-laki" {
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, profession, age)
	} else if gender == "perempuan" {
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", name, profession, age)
	}
}

// Soal 3
var buah = []string{}

func tambahBuah(data *[]string) {
	*data = append(*data, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
}

// Soal 4
var dataFilm = []map[string]string{}

func tambahDataFilm(title string, duration string, genre string, year string, data *[]map[string]string) {
	film := map[string]string{
		"Title":    title,
		"Duration": duration,
		"Genre":    genre,
		"Year":     year,
	}
	*data = append(*data, film)
}

func main() {
	// Soal 1
	hitungLingkaran(7.0, &luasLingkaran, &kelilingLingkaran)
	fmt.Println("Luas Lingkaran:", luasLingkaran)
	fmt.Println("Keliling Lingkaran:", kelilingLingkaran)

	// Soal 2
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	// Soal 3
	tambahBuah(&buah)
	for i, b := range buah {
		fmt.Printf("%d. %s\n", i+1, b)
	}

	// Soal 4
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s, Duration: %s, Genre: %s, Year: %s\n", i+1, film["Title"], film["Duration"], film["Genre"], film["Year"])
	}
}
