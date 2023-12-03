package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Digits int

const (
	One Digits = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

var digits = map[string]Digits{
	"one":   One,
	"two":   Two,
	"three": Three,
	"four":  Four,
	"five":  Five,
	"six":   Six,
	"seven": Seven,
	"eight": Eight,
	"nine":  Nine,
}

func main() {
	var finalTotal int
	f, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(f), "\n")

	for _, input := range inputs {
		first, last := getDigits(input)
		concantedString, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		log.Println(concantedString)
		finalTotal += concantedString
	}

	log.Println(finalTotal)
}

func getDigits(input string) (int, int) {
	var first int
	var firstDigitIndex int
	var last int
	var lastIndex int

	for i, c := range input {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			continue
		}

		if first == 0 {
			first = v
			firstDigitIndex = i
		}
		last = v
		lastIndex = i
	}

	for k := range digits {
		strIndex := strings.Index(input, k)
		lastStrIndex := strings.LastIndex(input, k)
		if strIndex != -1 {
			if strIndex <= firstDigitIndex {
				first = int(digits[k])
				firstDigitIndex = strIndex
			}
			if lastStrIndex >= lastIndex {
				last = int(digits[k])
				lastIndex = lastStrIndex
			}
		}
	}
	return first, last
}
