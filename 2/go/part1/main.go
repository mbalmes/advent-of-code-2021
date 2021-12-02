package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("../../input.txt")
	defer f.Close()

	var horizontal, depth int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		switch split[0] {
		case "down":
			x, _ := strconv.Atoi(split[1])
			depth += x
		case "up":
			x, _ := strconv.Atoi(split[1])
			depth -= x
		default:
			x, _ := strconv.Atoi(split[1])
			horizontal += x
		}
	}

	fmt.Println(horizontal * depth)
}
