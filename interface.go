package main

import (
	"fmt"
	"math"
)

// Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// Implementations for segitigaSamaSisi
func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

// Implementations for persegiPanjang
func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

// Implementations for tabung
func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

// Implementations for balok
func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64(b.panjang*b.lebar+b.panjang*b.tinggi+b.lebar*b.tinggi)
}

func main() {
	// Instances
	s := segitigaSamaSisi{alas: 3, tinggi: 4}
	p := persegiPanjang{panjang: 5, lebar: 6}
	t := tabung{jariJari: 7, tinggi: 8}
	b := balok{panjang: 9, lebar: 10, tinggi: 11}

	// Bangun Datar
	var bd hitungBangunDatar

	bd = s
	fmt.Println("Segitiga Sama Sisi")
	fmt.Println("Luas:", bd.luas())
	fmt.Println("Keliling:", bd.keliling())

	bd = p
	fmt.Println("\nPersegi Panjang")
	fmt.Println("Luas:", bd.luas())
	fmt.Println("Keliling:", bd.keliling())

	// Bangun Ruang
	var br hitungBangunRuang

	br = t
	fmt.Println("\nTabung")
	fmt.Println("Volume:", br.volume())
	fmt.Println("Luas Permukaan:", br.luasPermukaan())

	br = b
	fmt.Println("\nBalok")
	fmt.Println("Volume:", br.volume())
	fmt.Println("Luas Permukaan:", br.luasPermukaan())

	// Soal 2
	phoneStruct := phone{
		name:  "Samsung Galaxy Note 20",
		brand: "Samsung Galaxy Note 20",
		year:  2020,
		colors: []string{
			"Mystic Bronze", "Mystic White", "Mystic Black",
		},
	}

	var pi phoneInfo = phoneStruct
	fmt.Println("\nPhone Info")
	fmt.Println(pi.displayInfo())

	// Soal 3
	fmt.Println("\nLuas Persegi")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// Soal 4
	fmt.Println("\nHasil Penjumlahan")
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := 0
	for _, v := range angkaPertama {
		total += v
	}
	for _, v := range angkaKedua {
		total += v
	}

	fmt.Printf("%s %d + %d + %d + %d = %d", prefix, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)
}

// Soal 2 implementation
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phoneInfo interface {
	displayInfo() string
}

func (p phone) displayInfo() string {
	return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: %v", p.name, p.brand, p.year, p.colors)
}

// Soal 3 implementation
func luasPersegi(sisi int, verbose bool) interface{} {
	if sisi == 0 && verbose {
		return "Maaf anda belum menginput sisi dari persegi"
	}
	if sisi == 0 && !verbose {
		return nil
	}
	luas := sisi * sisi
	if verbose {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}
