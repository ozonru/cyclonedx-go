package bom

import (
	"os/exec"
	"bytes"
	"encoding/json"
	"log"
	"io"
	"strings"
	"encoding/xml"
	"github.com/google/uuid"
)

// TODO
// Required fields by spec:
// - type
// - name
// - version
type Module struct {
	XMLName xml.Name `xml:"component"`
	Type string `xml:"type,attr"`
	Path string `json:"Path"xml:"name"`
	main bool `json:"Main"`
	Version string `json:"Version"xml:"version"`
	indirect bool `json:"Indirect"`
}

func (m Module) NormalizeVersion(v string) string {
	return strings.TrimPrefix(v, "v")
}

// TODO See https://cyclonedx.org/docs/1.1/
type BOM struct {
	XMLName xml.Name `xml:"bom"`
	XMLNs string `xml:"xmlns,attr"`
	Version int `xml:"version,attr"`
	SerialNumber string `xml:"serialNumber,attr"`
	Modules []Module `xml:"components>component"`
}

// TODO
// (!) Semver conversion https://semver.org/#is-v123-a-semantic-version
func Generate() (string, error) {
	var result string
	cmd := exec.Command("go", "list", "-json", "-m", "all")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	bom := BOM{XMLNs:"http://cyclonedx.org/schema/bom/1.1", Version:1}
	bom.SerialNumber = uuid.New().URN()
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
		m.Version = m.NormalizeVersion(m.Version)
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
