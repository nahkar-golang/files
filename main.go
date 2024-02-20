package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "data.txt"
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString("Hello, World!\n")

	file_data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(string(file_data))

	os.Remove(fileName)
}
