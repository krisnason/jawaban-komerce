package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Membaca jumlah keluarga
	fmt.Print("Input the number of families: ")
	familiesStr, _ := reader.ReadString('\n')
	familiesStr = strings.TrimSpace(familiesStr)
	n, err := strconv.Atoi(familiesStr)
	if err != nil {
		fmt.Println("Jumlah keluarga tidak valid")
		return
	}

	// Membaca jumlah anggota di setiap keluarga
	fmt.Print("Input the number of members in the family ( separated by a space): ")
	membersStr, _ := reader.ReadString('\n')
	membersStr = strings.TrimSpace(membersStr)
	familySizesStr := strings.Split(membersStr, " ")
	if len(familySizesStr) != n {
		fmt.Println("Input must be equal with count of family")
		return
	}

	familySizes := make([]int, n)
	for i, s := range familySizesStr {
		familySizes[i], err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid number of family members input")
			return
		}
	}

	// Menghitung jumlah bus yang diperlukan
	jumlahBus := 0
	passengers := 0
	isFamilyCounted := make([]bool, n) // Menandai apakah sebuah keluarga sudah dihitung atau belum
	var familiesInBus []int
	var isFamilyCount []bool // Menandai apakah sebuah keluarga sudah dihitung atau belum

	// Loop through the family sizes
	for i := 0; i < n; i++ {
		if !isFamilyCounted[i] {
			if familySizes[i] == 4 {
				isFamilyCounted[i] = true
				jumlahBus++

			} else if familySizes[i] != 4 {
				// Cari keluarga lain yang total anggotanya 4
				for j := i + 1; j < n; j++ {
					if !isFamilyCounted[j] && familySizes[i]+familySizes[j] == 4 {
						isFamilyCounted[i] = true
						isFamilyCounted[j] = true
						jumlahBus++
						break
					}
				}
			}

			if !isFamilyCounted[i] {
				familiesInBus = append(familiesInBus, familySizes[i])
				isFamilyCount = append(isFamilyCount, false)
			}
		}
	}

	for i := 0; i < len(familiesInBus); i++ {

		// Skip the current family if it has been counted
		if !isFamilyCount[i] {
			// Check if current family can fit in the bus
			if passengers+familiesInBus[i] == 4 {
				// Jika jumlah anggota keluarga plus penumpang bus sekarang adalah 4 dan masih ada slot keluarga, maka keluarga ini dan keluarga lainnya masuk bus
				jumlahBus++
				passengers = 0
				isFamilyCount[i] = true
			} else if passengers+familiesInBus[i] < 4 {
				// Jika jumlah anggota keluarga plus penumpang bus sekarang kurang dari 4 dan masih ada slot keluarga, maka keluarga ini masuk bus
				passengers += familiesInBus[i]
				isFamilyCount[i] = true
			} else {
				// Jika bus sudah penuh atau tidak cukup tempat untuk keluarga ini, maka bus baru
				jumlahBus += (passengers + familiesInBus[i]) / 4
				passengers = (passengers + familiesInBus[i]) % 4
			}
		}

	}

	// If there are still remaining passengers, need another bus
	if passengers > 0 {
		jumlahBus++
	}

	fmt.Println(jumlahBus)
}
