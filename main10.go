package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

// soal 1
func printMessageAndYear(message string, year int) {
	fmt.Printf("%s - %d\n", message, year)
}

// soal 2
func kelilingSegitigaSamaSisi(sisi int, verbose bool) (string, error) {
	if sisi == 0 {
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		if verbose {
			return "", err
		} else {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()
			panic(err)
		}
	}

	keliling := sisi * 3
	if verbose {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}
	return fmt.Sprintf("%d", keliling), nil
}

// soal 3
func tambahAngka(n int, angka *int) {
	*angka += n
}

// Function untuk mencetak nilai akhir dari variabel angka
func cetakAngka(angka *int) {
	fmt.Printf("Total angka: %d\n", *angka)
}

// soal 4
func tambahPhone(phones *[]string, phone string) {
	*phones = append(*phones, phone)
}

// soal 5
func tampilkanFruitsWithGoroutine(fruits *[]string) {
	sort.Strings(*fruits)
	var wg sync.WaitGroup

	for i, fruit := range *fruits {
		wg.Add(1)
		go func(i int, phone string) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("%d. %s\n", i+1, phone)
		}(i, fruit)
	}
	wg.Wait()
}

// soal 6
func getMovies(moviesChannel chan string, movies ...string) {
	defer close(moviesChannel)
	for _, movie := range movies {
		moviesChannel <- movie
	}
}

func main() {
	// soal 1
	defer printMessageAndYear("Golang Backend Development", 2021)
	fmt.Printf("Fungsi printMessageAndYear akan dieksekusi terakhir.\n")

	// soal 2
	fmt.Println("Soal 2:")
	results := []struct {
		sisi    int
		verbose bool
	}{
		{4, true},
		{8, false},
		{0, true},
		{0, false},
	}

	for _, r := range results {
		result, err := kelilingSegitigaSamaSisi(r.sisi, r.verbose)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	}
	fmt.Println()

	// soal 3
	fmt.Println("Soal 3:")
	angka := 1

	// defer cetakAngka akan dieksekusi terakhir
	defer cetakAngka(&angka)

	// Tambah angka dengan beberapa nilai
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
	fmt.Println()

	// soal 4
	fmt.Println("Soal 4:")
	var phones []string

	tambahPhone(&phones, "Asus")
	tambahPhone(&phones, "Iphone")
	tambahPhone(&phones, "Oppo")
	tambahPhone(&phones, "Realme")
	tambahPhone(&phones, "Samsung")
	tambahPhone(&phones, "Vivo")
	tambahPhone(&phones, "Xiaomi")

	fmt.Println("Phones List:")
	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
	}
	fmt.Println()

	// soal 5
	fmt.Println("Soal 5:")
	fruits2 := []string{"Banana", "Apple", "Grape", "Mango", "Cherry", "Lemon", "Avocado"}

	fmt.Println("Phones List with Goroutine:")
	tampilkanFruitsWithGoroutine(&fruits2)
	fmt.Println()

	// soal 6
	fmt.Println("Soal 6:")
	movies := []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	fmt.Println("Movies List:")
	for value := range moviesChannel {
		fmt.Println(value)
	}
}
