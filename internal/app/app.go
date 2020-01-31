package app

import (
	"os/exec"
	"bytes"
	"encoding/json"
	"log"
	"io"
	"encoding/xml"
)

type Module struct {
	XMLName xml.Name `xml:"component"`
	Type string `xml:"type,attr"`
	Path string `json:"Path"xml:"name"`
	main bool `json:"Main"`
	Version string `json:"Version"xml:"version"`
	indirect bool `json:"Indirect"`
}

// TODO See https://cyclonedx.org/docs/1.1/
// <bom xmlns="http://cyclonedx.org/schema/bom/1.1" version="1" serialNumber="urn:uuid:deb8e876-42f7-4218-9661-db195cece217">
type BOM struct {
	XMLName xml.Name `xml:"bom"`
	Modules []Module `xml:"components>component"`
}

func Generate(outputFileName string) (string, error) {
	var result string
	cmd := exec.Command("go", "list", "-json", "-m", "all")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	bom := BOM{}
	dec := json.NewDecoder(bytes.NewReader(out))
	var modules []Module

	for {
		var m Module
		if err := dec.Decode(&m); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("reading go list output: %v", err)
		}
		m.Type = "library"
		modules = append(modules, m)
	}
	bom.Modules = modules
	xmlOut, err := xml.MarshalIndent(bom, " ", "  ")
	if err != nil {
		panic(err)
	}
	result = xml.Header + string(xmlOut)
	return result, nil
}
