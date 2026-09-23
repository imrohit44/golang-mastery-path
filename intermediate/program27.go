package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	ID      int      `xml:"id,attr"` // XML Attribute
	Name    string   `xml:"name"`    // XML Element
	Age     int      `xml:"age"`
}

func main() {
	xmlData := `
		<person id="101">
			<name>Rohit</name>
			<age>22</age>
		</person>`

	var p Person
	err := xml.Unmarshal([]byte(xmlData), &p)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}

	fmt.Printf("Parsed Struct: %+v\n", p)

	// Convert back to XML
	output, _ := xml.MarshalIndent(p, "", "  ")
	fmt.Println("\nGenerated XML:\n", string(output))
}