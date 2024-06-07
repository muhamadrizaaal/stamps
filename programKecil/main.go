package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk mengecek apakah sebuah angka adalah bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Buat array dari 1 sampai 100
	nums := make([]int, 100)
	for i := 0; i < 100; i++ {
		nums[i] = i + 1
	}

	// List untuk menampung hasil
	var result []string

	// Iterasi melalui array dari belakang ke depan
	for i := 99; i >= 0; i-- {
		num := nums[i]

		// Jika angka adalah bilangan prima, lewati
		if isPrime(num) {
			continue
		}

		// Ganti angka sesuai dengan aturan
		if num%3 == 0 && num%5 == 0 {
			result = append(result, "FooBar")
		} else if num%3 == 0 {
			result = append(result, "Foo")
		} else if num%5 == 0 {
			result = append(result, "Bar")
		} else {
			result = append(result, fmt.Sprintf("%d", num))
		}
	}

	// Print hasilnya dalam satu baris
	fmt.Println(strings.Join(result, " "))
}
