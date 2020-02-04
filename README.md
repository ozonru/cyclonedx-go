# cyclonedx-go

The CycloneDX module for Go creates a valid CycloneDX bill-of-material document containing an aggregate of all project dependencies. [CycloneDX](https://cyclonedx.org) is a lightweight BOM specification that is easily created, human readable, and simple to parse.

## Requirements

:warning: It works for projects with [Modules](https://blog.golang.org/using-go-modules) feauture enabled

## Install

### Local Installation

```bash
go get github.com/ozonru/cyclonedx-go/cmd/cyclonedx-go
```

## Usage

Navigate to the project directory and run `cyclonedx-go`. Inside it will read output from `go list -json -m all` command and print result BOM. You can specify destation for result file with option `-o`.

### Sample output

```bash
$ cyclonedx-go 
<?xml version="1.0" encoding="UTF-8"?>
 <bom xmlns="http://cyclonedx.org/schema/bom/1.1" version="1" serialNumber="urn:uuid:a83b2169-1ac1-4624-85ff-b5a4fa2fdd29">
   <components>
     <component type="library">
       <name>github.com/google/uuid</name>
       <version>1.1.1</version>
     </component>
   </components>
 </bom>
```

## License

Permission to modify and redistribute is granted under the terms of the GPL-3 license.
