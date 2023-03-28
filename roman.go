package roman

import (
	"math"
	"strings"
)

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
	"N": 0,
}

// romanToInt converts roman numeral to integer.
func romanToInt(romanNumeral string) int {
	if romanNumeral == "" {
		return 0
	}
	total, lastnum, num := 0, 0, 0
	for i := 0; i < len(romanNumeral); i++ {
		char := romanNumeral[len(romanNumeral)-(i+1) : len(romanNumeral)-i]
		num = roman[char]

		if num < lastnum {
			total = total - num
		} else {
			total = total + num
		}

		lastnum = num

	}

	return total
}

// vinculumRomanToInt converts a vinculum type of roman numberal (e.g where adding a horizontal line over a number multiplies it by 1000) to integer.
func vinculumRomanToInt(romanNumeral string) int {

	blocks := strings.Split(romanNumeral, "|")

	//coefficient for multiplying
	q := len(blocks) - 1

	total := 0

	for _, block := range blocks {
		num := romanToInt(block)
		total += num * int(math.Pow(float64(1000), float64(q)))
		q--
	}

	return total
}

// CompareTwoRomans compares two roman numberals.
func compareTwoRomans(roman1, roman2 string) int {

	if vinculumRomanToInt(roman1) == vinculumRomanToInt(roman2) {
		return 1
	}

	if vinculumRomanToInt(roman1) > vinculumRomanToInt(roman2) {
		return 1
	} else {
		return -1
	}

}
