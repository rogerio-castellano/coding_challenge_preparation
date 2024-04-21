package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Function to read a file and return its content
func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Function to write a string to a file
func writeFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// Function to append a string to an existing file
func appendToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// Function to demonstrate reading from a file line by line
func readLines(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}

func main() {
	// Example file names
	inputFile := "input.txt"
	outputFile := "output.txt"

	// Read from a file
	content, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", content)

	// Write to a file
	err = writeFile(outputFile, "This is a sample output.")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Append to a file
	err = appendToFile(outputFile, "\nAppending a second line.")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	// Read lines from a file
	err = readLines(inputFile)
	if err != nil {
		fmt.Println("Error reading lines from file:", err)
	}
}
