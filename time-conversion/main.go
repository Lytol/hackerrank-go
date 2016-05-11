package main

import (
	"fmt"
	"log"
	"time"
)

const inputLayout = "03:04:05PM"
const outputLayout = "15:04:05"

func main() {
	var str string

	_, err := fmt.Scanln(&str)
	if err != nil {
		log.Fatal(err)
	}

	t, err := time.Parse(inputLayout, str)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.Format(outputLayout))
}
