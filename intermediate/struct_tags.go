package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name" db:"first_name" xml:"first_name"`
	LastName  string `json:"last_name,omitempty" db:"last_name" xml:"last_name"`
	Age       int    `json:"-" db:"-" xml:"-"`
}

func main() {
	person := Person{
		FirstName: "Aishwarya",
		LastName:  "Rai",
		Age:       40,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
