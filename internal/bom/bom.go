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

type Module struct {
	Path string `json:"Path"`
	Main bool `json:"Main"`
	Version string `json:"Version"`
	Indirect bool `json:"Indirect"`
}

type Component struct {
	XMLName xml.Name `xml:"component"`
	Type string `xml:"type,attr"`
	Name string `xml:"name"`
	Version string `xml:"version"`
}

func (m Module) NormalizeVersion(v string) string {
	return strings.TrimPrefix(v, "v")
}

// See https://cyclonedx.org/docs/1.1/
type BOM struct {
	XMLName xml.Name `xml:"bom"`
	XMLNs string `xml:"xmlns,attr"`
	Version int `xml:"version,attr"`
	SerialNumber string `xml:"serialNumber,attr"`
	Components []Component `xml:"components>component"`
}

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
	var components []Component

	for {
		var m Module
		var c Component
		if err := dec.Decode(&m); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("reading go list output: %v", err)
		}
		if m.Main != true {
			c.Name = m.Path
			c.Type = "library"
			c.Version = m.NormalizeVersion(m.Version)
			components = append(components, c)
		}
	}
	bom.Components = components
	xmlOut, err := xml.MarshalIndent(bom, " ", "  ")
	if err != nil {
		panic(err)
	}
	result = xml.Header + string(xmlOut)
	return result, nil
}