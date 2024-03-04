package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findPair(arr []int, m int) (int, int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] > m {
				return arr[i], arr[j]
			}
		}
	}
	return -1, -1
}

func main() {
	// get array numbers from user
	fmt.Print("Enter an array of numbers (comma-separated): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputArray := scanner.Text()

	// convert string to number
	stringNumbers := strings.Split(inputArray, ",")
	var numbers []int
	for _, strNum := range stringNumbers {
		num, err := strconv.Atoi(strings.TrimSpace(strNum))
		if err != nil {
			fmt.Println("Invalid input. Please enter valid numbers.")
			return
		}
		numbers = append(numbers, num)
	}

	// get m number from user
	fmt.Print("Enter a number (m): ")
	var m int
	_, err := fmt.Scan(&m)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number for m.")
		return
	}

	// find pair numbers
	resultA, resultB := findPair(numbers, m)

	// show resultes
	fmt.Printf("Pair found: (%d, %d)\n", resultA, resultB)
}
