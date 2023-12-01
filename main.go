package main

import (
	"aoc2023/cmd"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getInput(day int) error {
	// Parse cookie
	cookieBuf, err := os.ReadFile(".cookie")
	if err != nil {
		return err
	}
	cookie := strings.TrimSuffix(string(cookieBuf), "\n")

	// Setup request
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day),
		nil,
	)
	if err != nil {
		return err
	}
	// Set cookie
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", string(cookie)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// Read response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Write to file
	err = os.WriteFile("input.txt", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// func main() {
// 	fmt.Println("Advent of code 2023")

// 	// Get input of the day
// 	err := getInput(1)
// 	if err != nil {
// 		fmt.Println("Failed to fetch input")
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	fmt.Println("Advent of code 2023")
// }

func main() {
	cmd.Execute()
}
