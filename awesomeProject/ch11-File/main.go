package main

import (
	"fmt"
	"os"
)

func main() {
	// Open a file
	file, err := os.Open("ch11-File/test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	fmt.Printf("File %v opened successfully", file.Name())

	// Read the file
	content, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(content))

	// Close the file
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
		return
	}
	fmt.Println("File closed successfully")
}
