package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/ozonru/cyclonedx-go/internal/app"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var outputFileName string

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Generate software bill-of-material (SBOM) file for Go project.\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of program:\n")
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nFields marked with (*) are required.\n")
	}

	// TODO
	// 1. Check if Go binary is installed and its version
	// 2. Check for go.mod
	// 3. Get result and print or write it to the file

	flag.StringVar(&outputFileName, "o", "bom.xml", "Result SBOM file")
	flag.Parse()
	result, err := app.Generate(outputFileName)
	checkError(err)
	fmt.Println(result)
}
