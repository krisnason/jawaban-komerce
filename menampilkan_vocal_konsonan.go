package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Menerima input dari pengguna
	fmt.Println("Masukkan string:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// input := "Sample Case"

	// Menghapus karakter newline dari input
	input = strings.TrimSpace(input)

	// Mengonversi string input ke huruf kecil agar penghitungan tidak case-sensitive
	input = strings.ToLower(input)

	// Map untuk menyimpan indeks kemunculan masing-masing karakter vokal dan konsonan
	vowelIndex := make(map[rune][]int)
	consonantIndex := make(map[rune][]int)

	// Melakukan iterasi pada setiap karakter dalam string input
	for i, char := range input {
		// Memeriksa apakah karakter adalah huruf
		if char >= 'a' && char <= 'z' {
			// Memeriksa apakah karakter adalah vokal atau konsonan
			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
				// Menyimpan indeks kemunculan karakter vokal
				vowelIndex[char] = append(vowelIndex[char], i)
			} else {
				// Menyimpan indeks kemunculan karakter konsonan
				consonantIndex[char] = append(consonantIndex[char], i)
			}
		}
	}

	// Menampilkan karakter vokal sesuai urutan kemunculannya
	fmt.Println("Karakter Vokal:")
	printCharactersInOrder(vowelIndex, input)

	// Menampilkan karakter konsonan sesuai urutan kemunculannya
	fmt.Println("\nKarakter Konsonan:")
	printCharactersInOrder(consonantIndex, input)
}

// Fungsi untuk mencetak karakter berdasarkan urutan kemunculannya
func printCharactersInOrder(index map[rune][]int, input string) {
	// Iterasi melalui setiap karakter dalam urutan kemunculan
	for _, char := range input {
		// Cek apakah karakter terdapat dalam map indeks
		if indices, ok := index[char]; ok {
			// Jika terdapat, cetak karakter sesuai dengan indeksnya tanpa spasi
			for _, idx := range indices {
				fmt.Printf("%c", input[idx])
			}
			// Setelah mencetak, hapus karakter dari map untuk menghindari pencetakan ganda
			delete(index, char)
		}
	}
}
