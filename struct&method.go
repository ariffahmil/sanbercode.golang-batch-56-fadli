package main

import "fmt"

// Soal 1
type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

// Soal 2
type segitiga struct {
	alas, tinggi int
}

func (s segitiga) luas() int {
	return (s.alas * s.tinggi) / 2
}

type persegi struct {
	sisi int
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

type persegiPanjang struct {
	panjang, lebar int
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

// Soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) tambahWarna(warna string) {
	p.colors = append(p.colors, warna)
}

// Soal 4
type movie struct {
	title    string
	duration string
	genre    string
	year     int
}

func tambahDataFilm(title string, duration string, genre string, year int, dataFilm *[]movie) {
	*dataFilm = append(*dataFilm, movie{title: title, duration: duration, genre: genre, year: year})
}

func main() {
	// Soal 1
	nanas := buah{nama: "Nanas", warna: "Kuning", adaBijinya: false, harga: 9000}
	jeruk := buah{nama: "Jeruk", warna: "Oranye", adaBijinya: true, harga: 8000}
	semangka := buah{nama: "Semangka", warna: "Hijau & Merah", adaBijinya: true, harga: 10000}
	pisang := buah{nama: "Pisang", warna: "Kuning", adaBijinya: false, harga: 5000}

	fmt.Println(nanas)
	fmt.Println(jeruk)
	fmt.Println(semangka)
	fmt.Println(pisang)

	// Soal 2
	s := segitiga{alas: 10, tinggi: 5}
	p := persegi{sisi: 4}
	pp := persegiPanjang{panjang: 7, lebar: 3}

	fmt.Println("Luas segitiga:", s.luas())
	fmt.Println("Luas persegi:", p.luas())
	fmt.Println("Luas persegi panjang:", pp.luas())

	// Soal 3
	ph := phone{name: "iPhone", brand: "Apple", year: 2020, colors: []string{"Black", "White"}}

	ph.tambahWarna("Red")
	ph.tambahWarna("Blue")

	fmt.Println(ph)

	// Soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", "2 jam", "action", 1999, &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", 2004, &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", 2004, &dataFilm)

	for _, film := range dataFilm {
		fmt.Println(film)
	}
}
