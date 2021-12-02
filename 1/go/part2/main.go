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

	var window [4]int
	for i := 1; i < 4; i++ {
		scanner.Scan()
		window[i], _ = strconv.Atoi(scanner.Text())
	}

	count := 0
	for scanner.Scan() {
		window[0], window[1], window[2] = window[1], window[2], window[3]
		window[3], _ = strconv.Atoi(scanner.Text())
		if window[0]+window[1]+window[2] < window[1]+window[2]+window[3] {
			count++
		}
	}

	fmt.Println(count)
}
