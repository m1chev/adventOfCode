package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1 - A Rock
// 2 - B Paper
// 3 - C Scissors
func result(line string) int {
	switch line{
	case "A X":		//Loos
		return 0+3
	case "B X":
		return 0+1
	case "C X":
		return 0+2
	case "B Y":		//Draw
		return 3+2
	case "A Y":
		return 3+1
	case "C Y":
		return 3+3
	case "C Z":		//Win
		return 6+1
	case "A Z":
		return 6+2
	case "B Z":
		return 6+3
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
