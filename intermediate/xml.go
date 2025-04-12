package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"first_name"`
	LastName  string   `xml:"last_name,omitempty"`
	Age       int      `xml:"-"`
	Address   Address  `xml:"address"`
}

type Address struct {
	Street  string   `xml:"street"`
	City    string   `xml:"city"`
	State   string   `xml:"state"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
	ISBN    string   `xml:"isbn,attr"`
	Pages   Pages    `xml:"pages"`
}

type Pages struct {
	PagesTotal   int      `xml:"pages_total,attr"`
	ColorPages int    `xml:"color_pages,attr"`
}

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
			State:  "CA",
		},
	}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))

	fmt.Println("--------------------------------")

	xmlData, err = xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))

	fmt.Println("--------------------------------")

	xmlData = []byte(`<actress><first_name>Rashmika</first_name><last_name>Mandanna</last_name><dimensions><height>5'6</height><weight>50 kg</weight></dimensions></actress>`)

	type Dimensions struct {
		Height string `xml:"height"`
		Weight string `xml:"weight"`
	}

	type Actress struct {
		XMLName xml.Name `xml:"actress"`
		FirstName string   `xml:"first_name"`
		LastName  string   `xml:"last_name"`
		Dimensions  Dimensions `xml:"dimensions"`
	}

	var decodedActress Actress
	err = xml.Unmarshal(xmlData, &decodedActress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(decodedActress)
	fmt.Println(decodedActress.XMLName.Local)
	fmt.Println(decodedActress.XMLName.Space)

	fmt.Println("--------------------------------")

	book := Book{
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		ISBN: "978-0-7432-7356-5",
		Pages: Pages{
			PagesTotal: 200,
			ColorPages: 50,
		},
	}
	xmlData, err = xml.MarshalIndent(book, "k", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))
	
	


	
}
