package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// part 02
func wordToNumber(word string) string {
	resultWord := word
	var tmpResult string
	listNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, digit := range listNumbers {
		tmpResult = resultWord
		if strings.Contains(word, digit) {
			if digit == "one" {
				resultWord = strings.Replace(tmpResult, digit, fmt.Sprintf("%s%s%s", string(digit[0]), "1", string(digit[len(digit)-1:])), -1)
			}
			if digit == "two" {
				resultWord = strings.Replace(tmpResult, digit, fmt.Sprintf("%s%s%s", string(digit[0]), "2", string(digit[len(digit)-1:])), -1)
			}
			if digit == "three" {
				resultWord = strings.Replace(tmpResult, digit, fmt.Sprintf("%s%s%s", string(digit[0]), "3", string(digit[len(digit)-1:])), -1)
			}
			if digit == "four" {
				resultWord = strings.Replace(tmpResult, digit, fmt.Sprintf("%s%s%s", string(digit[0]), "4", string(digit[len(digit)-1:])), -1)
			}
			if digit == "five" {
				resultWord = strings.Replace(tmpResult, digit, fmt.Sprintf("%s%s%s", string(digit[0]), "5", string(digit[len(digit)-1:])), -1)
			}
			if digit == "six" {
				resultWord = strings.Replace(tmpResult, digit, fmt.Sprintf("%s%s%s", string(digit[0]), "6", string(digit[len(digit)-1:])), -1)
			}
			if digit == "seven" {
				resultWord = strings.Replace(tmpResult, digit, fmt.Sprintf("%s%s%s", string(digit[0]), "7", string(digit[len(digit)-1:])), -1)
			}
			if digit == "eight" {
				resultWord = strings.Replace(tmpResult, digit, fmt.Sprintf("%s%s%s", string(digit[0]), "8", string(digit[len(digit)-1:])), -1)
			}
			if digit == "nine" {
				resultWord = strings.Replace(tmpResult, digit, fmt.Sprintf("%s%s%s", string(digit[0]), "9", string(digit[len(digit)-1:])), -1)
			}
		}
	}
	return resultWord
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// part 01
func isNumeric(word string) string {
	regexRule := regexp.MustCompile(`^[0-9]+$`)
	var result string
	for _, c := range word {
		numeric := regexRule.MatchString(string(c))
		if numeric {
			result = result + string(c)
		}
	}

	if len(result) > 2 {
		return string(result[0]) + string(result[len(result)-1:])
	}

	if len(result) == 1 {
		return result + result
	}

	return result
}

func main() {
	sumAll := 0
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		textoToNumber := wordToNumber(fileScanner.Text())
		extractNumber := isNumeric(textoToNumber)

		if len(extractNumber) == 1 {
			extractNumber = extractNumber + extractNumber
		}

		fmt.Printf("Original string: %s - Numbers: %s\n", fileScanner.Text(), extractNumber)
		sum, _ := strconv.Atoi(extractNumber)
		sumAll += sum
	}

	fmt.Printf("Total sum: %d\n", sumAll)

	readFile.Close()
}
