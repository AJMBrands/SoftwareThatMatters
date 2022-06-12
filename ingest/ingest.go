package ingest

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Metadata struct {
	XMLName    xml.Name   `xml:"metadata"`
	GroupId    string     `xml:"groupid"`
	ArtifactId string     `xml:"artifactid"`
	Versioning Versioning `xml:"versioning"`
}

type Versioning struct {
	XMLName     xml.Name `xml:"versioning"`
	Latest      string   `xml:"latest"`
	Release     string   `xml:"release"`
	Versions    []string `xml:"versions"`
	LastUpdated string   `xml:"lastupdated"`
}

func IngestData() {

	xmlFile, err := os.Open("maven-metadata.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	rawDataAddr, _ := ioutil.ReadAll(xmlFile)

	var parsed Metadata

	xml.Unmarshal(rawDataAddr, &parsed)

	version_number := parsed.Versioning.Versions
	fmt.Println(version_number)
}
