package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"golang.org/x/net/html"
)

const (
	URL = "https://adventofcode.com/2023/day"
)

func initDay(day int) error {

	fmt.Printf("Initializing day %d...\n", day)

	// Initializing directory
	dirName := fmt.Sprintf("day%02d", day)
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		fmt.Printf("Creating folder: %s/\n", dirName)
		if err := os.Mkdir(dirName, os.ModePerm); err != nil {
			return err
		}
	}

	// Fetch description
	err = fetchDescription(dirName, day)

	// Fetch input
	err = fetchInput(dirName, day)

	if err != nil {
		return err
	}

	return nil

}

func fetchDescription(dirName string, day int) error {

	url := fmt.Sprintf("%s/%d", URL, day)
	fmt.Println("Getting description from:", url)

	body, err := getHtmlBody(url)
	if err != nil {
		return err
	}
	bodyStr := string(body)

	// Parse html
	doc, err := html.Parse(strings.NewReader(bodyStr))
	if err != nil {
		return err
	}

	// Find the <article> tag
	node := make([]*html.Node, 0)

	var findArticle func(*html.Node)
	findArticle = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "article" {
			node = append(node, n)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findArticle(c)
		}
	}
	findArticle(doc)

	if node == nil {
		return fmt.Errorf("Article tag not found")
	}

	converter := md.NewConverter("", true, nil)

	desc := make([]string, 0)
	for _, n := range node {
		var src strings.Builder
		html.Render(&src, n)
		if err != nil {
			return err
		}
		text, err := converter.ConvertString(src.String())
		if err != nil {
			return err
		}
		desc = append(desc, text)
	}

	// Write to file
	fileName := filepath.Join(dirName, "README.md")
	fmt.Printf("Writing description: %s\n", fileName)
	err = os.WriteFile(fileName, []byte(strings.Join(desc, "\n\n")), 0644)
	if err != nil {
		return err
	}

	return nil

}

func getHtmlBody(url string) ([]byte, error) {
	// Parse cookie
	cookieBuf, err := os.ReadFile(".cookie")
	if err != nil {
		return nil, err
	}
	cookie := strings.TrimSuffix(string(cookieBuf), "\n")

	// Setup request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// Set cookie
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: string(cookie),
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Error %s: %s", resp.Status, body)
	}
	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func fetchInput(dirName string, day int) error {
	fileName := filepath.Join(dirName, "input.txt")

	// Check if input already exists
	_, err := os.Stat(fileName)
	if !os.IsNotExist(err) {
		return nil
	}

	url := fmt.Sprintf("%s/%d/input", URL, day)

	body, err := getHtmlBody(url)
	if err != nil {
		return err
	}

	// Write to file
	fmt.Printf("Writing input: %s\n", fileName)
	err = os.WriteFile(fileName, body, 0644)
	if err != nil {
		return err
	}

	return nil
}
