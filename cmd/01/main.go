package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Initate finalTotal variable
	var finalTotal int

	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")

	for _, l := range input {
		var first string
		var second string

		splitLine := strings.Split(l, "")

		// Check each char in word and find first and last digit
		for i := 0; i < len(splitLine); i++ {
			_, err := strconv.Atoi(splitLine[i])
			if err != nil {
				continue
			}
			if first == "" {
				first = splitLine[i]
			}
			second = splitLine[i]
		}

		concatVal, _ := strconv.Atoi(first + second)
		finalTotal += concatVal
	}
	log.Println(finalTotal)
}
