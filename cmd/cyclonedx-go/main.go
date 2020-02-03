package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/ozonru/cyclonedx-go/internal/bom"
	"io/ioutil"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var outputFileName string

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Generate software bill-of-material (SBOM) file for Go project (with modules).\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of program:\n")
		flag.PrintDefaults()
	}

	// TODO
	// 1. Check if Go binary is installed and its version
	flag.StringVar(&outputFileName, "o", "", "Result SBOM file")
	flag.Parse()

	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println("Can't find go.mod file in the current working directory.");
		os.Exit(1)
	}

	result, err := bom.Generate()
	checkError(err)

	if outputFileName != "" {
		err := ioutil.WriteFile(outputFileName, []byte(result), 0644)
		checkError(err)
	} else {
		fmt.Println(result)
	}
}
