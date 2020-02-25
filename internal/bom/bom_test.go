package bom

import (
	"testing"
	"encoding/xml"
)

var inputData []byte = []byte(`{
"Path": "github.com/google/uuid",
"Version": "v1.1.1",
"Time": "2019-02-27T21:05:49Z",
"Dir": "/go/pkg/mod/github.com/google/uuid@v1.1.1",
"GoMod": "/go/pkg/mod/cache/download/github.com/google/uuid/@v/v1.1.1.mod"
}`)

func TestName(t *testing.T) {
	var b BOM
	want := "github.com/google/uuid"
	xmlResult, _ := GenerateFromJSON(inputData)
	_ = xml.Unmarshal([]byte(xmlResult), &b)
	if got := b.Components[0].Name; got != want {
		t.Errorf(
		"Package name from result CycloneDX BOM = %q, want %q",
		got,
		want,
		)
	}
}

func TestVersion(t *testing.T) {
	var b BOM
	want := "1.1.1"
	xmlResult, _ := GenerateFromJSON(inputData)
	_ = xml.Unmarshal([]byte(xmlResult), &b)
	if got := b.Components[0].Version; got != want {
		t.Errorf(
		"Package version from result CycloneDX BOM = %q, want %q",
		got,
		want,
		)
	}
}
