package app

import (
	"os/exec"
)

func Generate(outputFileName string) (string, error) {
	// Algo
	// 1. Get output from go list
	// 2. Parse JSON
	// 3. Make XML
	var result string
	cmd := exec.Command("go", "list", "-json", "-m", "all")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	result = string(out)
	return result, nil
}
