package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

const (
	URL = "https://adventofcode.com/2023/day"
)

func initDay(day int) error {

	fmt.Printf("Initializing day %d...\n", day)

	// Initializing directory
	dirName := fmt.Sprintf("day%02d", day)
	// _, err := os.Stat(dirName)
	// if !os.IsNotExist(err) {
	// 	return fmt.Errorf("Day %d already initialized!\n", day)
	// }
	// fmt.Printf("Creating folder: %s/\n", dirName)
	// if err := os.Mkdir(dirName, os.ModePerm); err != nil {
	// 	return err
	// }

	// Fetch description
	// err := fetchDescription(dirName, day)

	// Fetch input
	err := fetchInput(dirName, day)
	// fmt.Sprintf("%s/%d/input", URL, day)

	if err != nil {
		return err
	}

	return nil

}

func fetchDescription(dirName string, day int) error {

	url := fmt.Sprintf("%s/%d", URL, day)
	fmt.Println("Getting description from:", url)

	data, err := getHtmlData(url)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(data, &buf); err != nil {
		return err
	}

	// err = os.WriteFile("text.txt", []byte(buf.String()), 0644)
	fmt.Println(buf.String())
	return err

}

func getHtmlData(url string) ([]byte, error) {
	// Parse cookie
	cookieBuf, err := os.ReadFile(".cookie")
	if err != nil {
		return nil, err
	}
	cookie := strings.TrimSuffix(string(cookieBuf), "\n")

	// Setup request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Set cookie
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", string(cookie)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Read response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func fetchInput(dirName string, day int) error {
	url := fmt.Sprintf("%s/%d/input", URL, day)

	data, err := getHtmlData(url)
	if err != nil {
		return err
	}

	// Write to file
	fileName := filepath.Join(dirName, "input.txt")
	fmt.Printf("Writing input: %s\n", fileName)
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
