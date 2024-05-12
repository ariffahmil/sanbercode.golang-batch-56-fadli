package main

import "fmt"

func main() {
	// Soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("%d - Berkualitas\n", i)
		} else if i%3 == 0 && i%2 != 0 {
			fmt.Printf("%d - I Love Coding\n", i)
		} else {
			fmt.Printf("%d - Santai\n", i)
		}
	}

	// Soal 2
	fmt.Println()
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// Soal 3
	fmt.Println()
	kalimat := []string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	// Soal 4
	fmt.Println()
	sayuran := []string{
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	}
	for i, sayur := range sayuran {
		fmt.Printf("%d. %s\n", i+1, sayur)
	}

	// Soal 5
	fmt.Println()
	satuan := map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	for key, value := range satuan {
		fmt.Printf("%s: %d\n", key, value)
	}
}
