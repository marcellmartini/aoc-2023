package puzzle

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

func GetPuzzle(url string) {
	c := colly.NewCollector()

	var textExtracted string

	// Find and visit all links
	c.OnHTML("article.day-desc", func(e *colly.HTMLElement) {
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
	fmt.Println(textExtracted)
}

// extractText recursively traverses the HTML tree and extracts the text content
func extractText(n *html.Node) string {
	var result strings.Builder
	var processNode func(*html.Node)

	processNode = func(node *html.Node) {
		if node.Type == html.TextNode {
			result.WriteString(node.Data)
		} else if node.Type == html.ElementNode {
			if node.Data == "p" || node.Data == "h2" {
				// Add newline before paragraphs and headers
				if result.Len() > 0 {
					result.WriteString("\n")
				}
			}

			for c := node.FirstChild; c != nil; c = c.NextSibling {
				processNode(c)
			}
		} else {
			for c := node.FirstChild; c != nil; c = c.NextSibling {
				processNode(c)
			}
		}
	}

	// Start processing the root node
	processNode(n)

	return result.String()
}
