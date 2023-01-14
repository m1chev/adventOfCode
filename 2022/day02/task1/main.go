package main

import (
	"bufio"
	"fmt"
	"os"
)

func result(line string) int {
	switch line{
	case "A X": //Rock
		return 3+1
	case "A Y":
		return 6+2
	case "A Z":
		return 0+3
	case "B Y": //Paper
		return 3+2
	case "B X":
		return 0+1
	case "B Z":
		return 6+3
	case "C Z": //Scissors
		return 3+3
	case "C X":
		return 6+1
	case "C Y":
		return 0+2
	}
	return 0
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += result(line)
	}

	fmt.Println(sum)
}
