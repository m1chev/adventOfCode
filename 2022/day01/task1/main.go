package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var sum, max int

	for scanner.Scan() {
        line := scanner.Text()
        if line != "" {
            calories, err := strconv.Atoi(line)
            if err != nil {
                fmt.Print(err)
            }
            sum += calories
        } else {
            if max < sum {
                max = sum
            }
            sum = 0
        }
	}
    fmt.Println(max)
}
