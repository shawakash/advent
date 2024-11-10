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

	gameIDStr := strings.Fields(parts[0])[1]
	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		panic("Invalid game ID")
	}

	sets := strings.Split(parts[1], ";")
	for _, set := range sets {
		var redCount, greenCount, blueCount int

		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			cubeInfo := strings.Fields(cube)
			count, err := strconv.Atoi(cubeInfo[0])
			if err != nil {
				panic("Invalid count in cube information")
			}

			switch cubeInfo[1] {
			case "red":
				redCount = count
			case "green":
				greenCount = count
			case "blue":
				blueCount = count
			}
		}

		if redCount > red || greenCount > green || blueCount > blue {
			return 0
		}
	}

	return gameID
}
