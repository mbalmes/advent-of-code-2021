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

	var horizontal, depth, aim int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		switch split[0] {
		case "down":
			x, _ := strconv.Atoi(split[1])
			aim += x
		case "up":
			x, _ := strconv.Atoi(split[1])
			aim -= x
		default:
			x, _ := strconv.Atoi(split[1])
			horizontal += x
			depth += x * aim
		}
	}

	fmt.Println(horizontal * depth)
}
