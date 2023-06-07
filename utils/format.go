package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func FormatCode() {
	// code, err := ioutil.ReadFile("utils/index.js")
	// if err != nil {
	// 	panic(err)
	// }
	// // format the code

	// // Create an output file to write the highlighted code
	// outputFile, err := os.Create("highlighted.html")
	// if err != nil {
	// 	panic(err)
	// }
	// defer outputFile.Close()

	// // Use the Quick.Highlight method to perform syntax highlighting
	// err = quick.Highlight(outputFile, string(code), "go", "html", "monokai")
	// if err != nil {
	// 	panic(err)
	// }

	// // Highlighting is complete
	// fmt.Println("Code snippet highlighted successfully!")

	// Define the Go code you want to highlight
	code := `
		package main

		import "fmt"

		func main() {
			fmt.Println("Hello, World!")
		}
	`

	// Get a lexer for Go code
	lexer := lexers.Get("go")
	if lexer == nil {
		lexer = chroma.Coalesce(lexers.Analyse(code))
	}
	// Create a formatter with HTML output
	formatter := html.New(html.WithClasses(true))
	style := styles.Get("monokai")
	// Tokenize the code
	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		fmt.Println("Failed to tokenize code:", err)
		os.Exit(1)
	}

	// Create a buffer to save the HTML content
	buf := bytes.NewBuffer(nil)

	// Format the tokens and save the HTML content in the buffer
	err = formatter.Format(buf, style, iterator)
	if err != nil {
		fmt.Println("Failed to format code:", err)
		os.Exit(1)
	}

	// Get the generated HTML content as a string
	htmlContent := buf.String()

	// Save the HTML content to a file
	err = ioutil.WriteFile("generated.html", []byte(htmlContent), 0644)
	if err != nil {
		fmt.Println("Failed to save HTML content:", err)
		os.Exit(1)
	}

	fmt.Println("HTML content saved to generated.html")

}
