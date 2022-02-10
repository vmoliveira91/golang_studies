package main

import "fmt"

func romanToInt(s string) int {
	value := 0

	conversion := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			if i < (len(s)-1) && (s[i+1] == 'V' || s[i+1] == 'X') {
				value += conversion[string(s[i])+string(s[i+1])]
				i++
			} else {
				value += conversion[string(s[i])]
			}
		} else if s[i] == 'V' {
			value += conversion[string(s[i])]
		} else if s[i] == 'X' {
			if i < (len(s)-1) && (s[i+1] == 'L' || s[i+1] == 'C') {
				value += conversion[string(s[i])+string(s[i+1])]
				i++
			} else {
				value += conversion[string(s[i])]
			}
		} else if s[i] == 'L' {
			value += conversion[string(s[i])]
		} else if s[i] == 'C' {
			if i < (len(s)-1) && (s[i+1] == 'D' || s[i+1] == 'M') {
				value += conversion[string(s[i])+string(s[i+1])]
				i++
			} else {
				value += conversion[string(s[i])]
			}
		} else if s[i] == 'D' {
			value += conversion[string(s[i])]
		} else if s[i] == 'M' {
			value += conversion[string(s[i])]
		}
	}

	return value
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
