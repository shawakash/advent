package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Should have been able to read the file")
	}

	sum := 0
	scanner := bufio.NewScanner(strings.NewReader(string(contents)))
	for scanner.Scan() {
		line := scanner.Text()
		sum += processGame(line, 12, 13, 14)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

func processGame(line string, red, green, blue int) int {
	parts := strings.Split(line, ":")

	maxRed := 0
	maxBlue := 0
	maxGreen := 0

	sets := strings.Split(parts[1], ";")
	for _, set := range sets {

		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			cubeInfo := strings.Fields(cube)
			count, err := strconv.Atoi(cubeInfo[0])
			if err != nil {
				panic("Invalid count in cube information")
			}

			switch cubeInfo[1] {
			case "red":
				maxRed = max(maxRed, count)
			case "green":
				maxGreen = max(maxGreen, count)
			case "blue":
				maxBlue = max(maxBlue, count)
			}
		}
	}

	return maxRed * maxBlue * maxGreen
}
