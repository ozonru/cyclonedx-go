package bom

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/google/uuid"
	"github.com/package-url/packageurl-go"
	"io"
	"os/exec"
	"strings"
)

type Module struct {
	Path     string `json:"Path"`
	Main     bool   `json:"Main"`
	Version  string `json:"Version"`
	Indirect bool   `json:"Indirect"`
}

type Component struct {
	XMLName xml.Name `xml:"component"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Version string   `xml:"version"`
	PURL    string   `xml:"purl"`
}

func (m Module) NormalizeVersion(v string) string {
	return strings.TrimPrefix(v, "v")
}

func (m Module) PURL() string {
	var ns, n string
	n = m.Path
	chunks := strings.Split(m.Path, "/")

	if len(chunks) > 1 {
		ns = strings.Join(chunks[:len(chunks)-1], "/")
		n = chunks[len(chunks)-1]
	}

	p := packageurl.NewPackageURL(
		packageurl.TypeGolang,
		ns,
		n,
		m.NormalizeVersion(m.Version),
		nil,
		"")
	return p.ToString()
}

// See https://cyclonedx.org/docs/1.1/
type BOM struct {
	XMLName      xml.Name    `xml:"bom"`
	XMLNs        string      `xml:"xmlns,attr"`
	Version      int         `xml:"version,attr"`
	SerialNumber string      `xml:"serialNumber,attr"`
	Components   []Component `xml:"components>component"`
}

func Generate() (string, error) {
	var result string

	cmd := exec.Command("go", "list", "-json", "-m", "all")
	out, err := cmd.Output()

	if err != nil {
		return result, err
	}

	bom := BOM{XMLNs: "http://cyclonedx.org/schema/bom/1.1", Version: 1}
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
			return result, err
		}
		if m.Main != true {
			c.Name = m.Path
			c.Type = "library"
			c.PURL = m.PURL()
			c.Version = m.NormalizeVersion(m.Version)
			components = append(components, c)
		}
	}
	bom.Components = components
	xmlOut, err := xml.MarshalIndent(bom, " ", "  ")
	if err != nil {
		return result, nil
	}
	result = xml.Header + string(xmlOut)
	return result, nil
}
