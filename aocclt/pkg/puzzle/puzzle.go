package puzzle

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

func GetPuzzle(url string) {
	fmt.Println(getElement(url, "article.day-desc"))
}

func GetExample(url string) {
	fmt.Println(getElement(url, "pre"))
}

func getElement(url string, selector string) string {
	if url == "" {
		os.Exit(0)
	}

	c := colly.NewCollector()

	var textExtracted string

	// Find and visit all links
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		htmlContent, err := e.DOM.Html()
		if err != nil {
			log.Fatal(err)
		}

		// Parse the HTML content
		doc, err := html.Parse(strings.NewReader(htmlContent))
		if err != nil {
			log.Fatal(err)
		}

		// Extract the text from the parsed HTML
		textExtracted = extractText(doc)
	})

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping by visiting the URL containing the HTML
	if err := c.Visit(url); err != nil {
		log.Fatal(err)
	}

	// Print the extracted text with indentation
	return strings.TrimSpace(textExtracted)
}

// extractText recursively traverses the HTML tree and extracts the text content
func extractText(n *html.Node) string {
	var result strings.Builder

	// Start processing the root node
	processNode(&result, n)

	return result.String()
}

func processNode(r *strings.Builder, node *html.Node) {
	if node.Type == html.TextNode {
		r.WriteString(node.Data)
	} else if node.Type == html.ElementNode {
		if node.Data == "p" || node.Data == "h2" {
			// Add newline before paragraphs and headers
			if r.Len() > 0 {
				r.WriteString("\n")
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			processNode(r, c)
		}
	} else {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			processNode(r, c)
		}
	}
}
