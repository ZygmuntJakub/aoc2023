package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nums = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var nums_words = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	result1 := 0
	result2 := 0
	file, err := os.Open("inputs/1")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var numbers = ""
		var numbers2 = ""
		var word = ""
	main:
		for _, sign := range line {
			letter := string(sign)
			word += letter
			for k, v := range nums_words {
				if strings.Contains(word, k) {
					numbers2 += v
					word = ""
					continue main
				}
			}
			if _, exists := nums[letter]; exists {
				numbers += letter
				numbers2 += letter
			}
		}

		var sum, _ = strconv.Atoi(string(numbers[0]) + string(numbers[len(numbers)-1]))
		var sum2, _ = strconv.Atoi(string(numbers2[0]) + string(numbers2[len(numbers)-1]))
		result1 += sum
		result2 += sum2

	}

	fmt.Printf("Ex1: %v\n", result1)
	fmt.Printf("Ex2: %v\n", result2)

}
