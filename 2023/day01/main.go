package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func numberLiteral(s string) string {
	switch number := s; number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}

// part 02
func wordToNumber(word string, position string) string {
	var digitString string

	pattern := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)")
	matches := pattern.FindAllString(word, -1)

	if len(matches) > 0 {
		if position == "start" {
			digitString = numberLiteral(matches[0])
		}
		if position == "end" {
			digitString = numberLiteral(matches[len(matches)-1])
		}
		if position == "all" {
			for _, matche := range matches {
				digitString = digitString + numberLiteral(matche)
			}
		}
	} else {
		fmt.Println(matches)
	}

	return digitString

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
		if isNumber(fileScanner.Text()[len(fileScanner.Text())-1:]) && isNumber(string(fileScanner.Text()[0])) {
			result := string(fileScanner.Text()[0]) + fileScanner.Text()[len(fileScanner.Text())-1:]
			if len(result) == 1 {
				result = result + result
			}
			fmt.Printf("Original string (all number): %s - Numbers: %s\n", fileScanner.Text(), result)
			sum, _ := strconv.Atoi(result)
			sumAll += sum
			continue
		}

		if isNumber(fileScanner.Text()[len(fileScanner.Text())-1:]) {
			result := wordToNumber(fileScanner.Text(), "start")
			if len(result) == 1 {
				result = result + result
			}
			if len(result) > 2 {
				result = string(result[0]) + result[len(result)-1:]
			}
			fmt.Printf("Original string (second number): %s - Numbers: %s\n", fileScanner.Text(), string(result[0])+fileScanner.Text()[len(fileScanner.Text())-1:])
			sum, _ := strconv.Atoi(result)
			sumAll += sum

			continue
		}

		if !isNumber(fileScanner.Text()[len(fileScanner.Text())-1:]) && !isNumber(string(fileScanner.Text()[0])) {
			result := wordToNumber(fileScanner.Text(), "all")
			fmt.Printf("Original string (all text): %s - Numbers: %s\n", fileScanner.Text(), result)
			continue
		}

		// resultWord := wordToNumber(fileScanner.Text())
		// fmt.Printf("Original string: %s - Numbers: %s\n", fileScanner.Text(), resultWord)

		// number, err := strconv.Atoi(resultLine)
		// if err != nil {
		// 	panic(err)
		// }
		// sumAll = sumAll + number
	}

	fmt.Printf("Total sum: %d\n", sumAll)

	readFile.Close()
}
