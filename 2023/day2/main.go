package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkPossible(cubeCount int, cubeList []int) bool {
	for _, cube := range cubeList {
		if cube > cubeCount {
			return false
		}
	}
	return true
}

func extractCubeCount(line []string, cube string) []int {
	// sumAll := 0
	var result []string
	var resultInt []int
	re := regexp.MustCompile("[0-9]+")

	for _, v := range line {
		if strings.Contains(v, cube) {
			result = append(result, re.FindAllString(v, -1)...)
		}
	}

	for _, i := range result {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		resultInt = append(resultInt, j)
	}
	return resultInt
}

func main() {
	var setsNormalized []string
	redCube := 12
	greenCube := 13
	blueCube := 14
	sumGameIDs := 0

	fmt.Println("day 2\n")
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		gameID := strings.Split(strings.Split(fileScanner.Text(), ":")[0], " ")[1]

		for _, set := range strings.Split(strings.Split(fileScanner.Text(), ":")[1], ",") {
			if strings.Contains(set, ";") {
				setsNormalized = append(setsNormalized, strings.Split(set, ";")...)
			} else {
				setsNormalized = append(setsNormalized, set)
			}
		}
		// fmt.Printf("Game %s- Red: %d, Blue: %d, Green: %d\n", gameID, extractCubeCount(setsNormalized, "red"), extractCubeCount(setsNormalized, "green"), extractCubeCount(setsNormalized, "blue"))
		if checkPossible(redCube, extractCubeCount(setsNormalized, "red")) && checkPossible(greenCube, extractCubeCount(setsNormalized, "green")) && checkPossible(blueCube, extractCubeCount(setsNormalized, "blue")) {
			fmt.Printf("Game %s possible!\n, Red:%t, Green:%t, Blue:%t\n", gameID, checkPossible(redCube, extractCubeCount(setsNormalized, "red")), checkPossible(greenCube, extractCubeCount(setsNormalized, "green")), checkPossible(blueCube, extractCubeCount(setsNormalized, "blue")))
			sum, _ := strconv.Atoi(gameID)
			sumGameIDs += sum
		}

		// if (extractCubeCount(setsNormalized, "red") <= redCube) && (extractCubeCount(setsNormalized, "green") <= greenCube) && (extractCubeCount(setsNormalized, "blue") <= blueCube) {
		// 	fmt.Printf("Game %s possible! - Red: %d, Blue: %d, Green: %d\n", gameID, extractCubeCount(setsNormalized, "red"), extractCubeCount(setsNormalized, "green"), extractCubeCount(setsNormalized, "blue"))
		// 	sum, _ := strconv.Atoi(gameID)
		// 	sumGameIDs += sum
		// }
		setsNormalized = nil
	}

	fmt.Printf("Sum of possible games: %d\n", sumGameIDs)
}
