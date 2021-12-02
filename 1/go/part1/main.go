package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lastMeasurement := 170
	count := 0
	for scanner.Scan() {
		currentMeasurement, _ := strconv.Atoi(scanner.Text())
		if currentMeasurement > lastMeasurement {
			count++
		}
		lastMeasurement = currentMeasurement
	}

	fmt.Println(count)
}
