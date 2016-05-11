package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var height int

	_, err := fmt.Scanf("%d", &height)
	if err != nil {
		log.Fatal(err)
	}

	PrintStaircase(height)
}

func PrintStaircase(height int) {
	for i := 1; i <= height; i++ {
		fmt.Printf("%s%s\n",
			strings.Repeat(" ", height-i),
			strings.Repeat("#", i))
	}
}
