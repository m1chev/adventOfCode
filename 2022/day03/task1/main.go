package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var sum int

	for scanner.Scan() {
		line := scanner.Text()
		for _, i := range line[:len(line)/2] {
			for _, j := range line[len(line)/2:] {
				if i == j {
                    if i > 96 {
                        sum += int(i) - 96
                    } else {
                        sum += int(i) - 38
                    }
                    goto Next
				}
			}
		}
        Next:
	}
    fmt.Println(sum)
}
