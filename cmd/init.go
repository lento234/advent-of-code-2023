package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/charmbracelet/log"
	"golang.org/x/net/html"
)

const (
	URL = "https://adventofcode.com/2023/day"
)

func initDay(day int) error {

	log.Infof("Initializing day %d...", day)

	// Initializing directory
	dirName := fmt.Sprintf("day%02d", day)
	_, err := os.Stat(dirName)

	if os.IsNotExist(err) {

		log.Infof("Creating folder: '%s/'", dirName)
		if err := os.Mkdir(dirName, os.ModePerm); err != nil {
			return err
		}
		// Copy templates
		err = setupTemplates(dirName, day)
		if err != nil {
			return err
		}
		// Fetch input
		err = fetchInput(dirName, day)
		if err != nil {
			return err
		}
	}

	// Fetch description
	err = fetchDescription(dirName, day)
	if err != nil {
		return err
	}

	return nil

}

func copyTemplate(srcFilename, dstFilename string, params map[string]string) error {

	log.Infof("Creating template: %s -> %s", srcFilename, dstFilename)

	// Read file
	buf, err := os.ReadFile(srcFilename)
	if err != nil {
		return err
	}

	// Template
	tmpl := template.New(dstFilename)
	tmpl, err = tmpl.Parse(string(buf))
	if err != nil {
		return err
	}

	f, err := os.Create(dstFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write to file
	err = tmpl.Execute(f, params)
	if err != nil {
		return err
	}

	return nil

}

func setupTemplates(dirname string, day int) error {
	tmplDir := "template"

	params := map[string]string{
		"day": fmt.Sprintf("%02d", day),
	}
	// Copy solution
	err := copyTemplate(
		filepath.Join(tmplDir, "dayXX.go.tmpl"),
		filepath.Join(dirname, fmt.Sprintf("day%02d.go", day)),
		params,
	)
	if err != nil {
		return err
	}

	// Copy test
	err = copyTemplate(
		filepath.Join(tmplDir, "dayXX_test.go.tmpl"),
		filepath.Join(dirname, fmt.Sprintf("day%02d_test.go", day)),
		params,
	)
	if err != nil {
		return err
	}

	// Copy benchmark
	err = copyTemplate(
		filepath.Join(tmplDir, "benchmark_test.go.tmpl"),
		filepath.Join(dirname, "benchmark_test.go"),
		params,
	)
	if err != nil {
		return err
	}

	return nil

}

func fetchDescription(dirName string, day int) error {

	url := fmt.Sprintf("%s/%d", URL, day)
	log.Infof("Getting description from: '%s'", url)

	body, err := getHTMLBody(url)
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
		return fmt.Errorf("'article' tag not found")
	}

	converter := md.NewConverter("", true, nil)

	desc := make([]string, 0)
	for _, n := range node {
		var src strings.Builder
		err := html.Render(&src, n)
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
	log.Infof("Writing description: '%s'", fileName)
	err = os.WriteFile(fileName, []byte(strings.Join(desc, "\n\n")), 0644)
	if err != nil {
		return err
	}

	return nil

}

func getHTMLBody(url string) ([]byte, error) {
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
		Value: cookie,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error %s: %s", resp.Status, body)
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

	body, err := getHTMLBody(url)
	if err != nil {
		return err
	}

	// Write to file
	log.Infof("Writing input: %s", fileName)
	err = os.WriteFile(fileName, body, 0644)
	if err != nil {
		return err
	}

	return nil
}
