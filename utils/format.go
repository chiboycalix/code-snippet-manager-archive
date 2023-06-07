package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/chroma/quick"
)

func FormatCode() {
	code, err := ioutil.ReadFile("utils/index.js")
	if err != nil {
		panic(err)
	}
	// format the code

	// Create an output file to write the highlighted code
	outputFile, err := os.Create("highlighted.html")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Use the Quick.Highlight method to perform syntax highlighting
	err = quick.Highlight(outputFile, string(code), "go", "html", "monokai")
	if err != nil {
		panic(err)
	}

	// Highlighting is complete
	fmt.Println("Code snippet highlighted successfully!")
}
