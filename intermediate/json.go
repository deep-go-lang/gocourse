package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"f_name,omitempty"`
	LastName  string `json:"l_name,omitempty"`
	Age       int    `json:"age,omitempty"`
	Email     string `json:"email,omitempty"`
	Address   Address `json:"address,omitempty"`
}

type Address struct {
	Street string `json:"street,omitempty"`
	City   string `json:"city,omitempty"`
	State  string `json:"state,omitempty"`
}

func main() {
	listOfAddress := []Address{
		{
			Street: "peddar road",
			City:   "Mumbai",
			State:  "Maharashtra",
		},
		{
			Street: "Bandra",
			City:   "Mumbai",
			State:  "Maharashtra",
		},
		{
			Street: "Khar",
			City:   "Mumbai",
			State:  "Maharashtra",
		},
	}

	actress := Person{
		FirstName: "Kriti",
		LastName:  "Sanon",
		Email:     "kriti@gmail.com",
		Address: Address{
			Street: "peddar road",
			City:   "Mumbai",
			State:  "Maharashtra",
		},
	}

	jsonData, err := json.Marshal(actress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	fmt.Println("--------------------------------")

	jsonData, err = json.MarshalIndent(actress, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	fmt.Println("--------------------------------")

	actressData := `{"first_name":"Kriti","last_name":"Sanon","age":28,"height":"5'9","address":{"street":"peddar road","city":"Mumbai","state":"Maharashtra"}}`

	type MyActress struct {
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Age       int    `json:"age,omitempty"`
		Height    string `json:"height,omitempty"`
		Address   Address `json:"address,omitempty"`
	}

	var actress3 MyActress

	err = json.Unmarshal([]byte(actressData), &actress3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("actress full name is", actress3.FirstName, actress3.LastName)
	fmt.Println("actress age next year will be", actress3.Age+1)
	fmt.Println("actress height is", actress3.Height)
	fmt.Println("actress address is", actress3.Address.Street, actress3.Address.City, actress3.Address.State)

	fmt.Println("--------------------------------")

	fmt.Println(listOfAddress)

	jsonData, err = json.Marshal(listOfAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	fmt.Println("--------------------------------")

	modelJson := `{"model":"Tapsee Pannu","age":32,"address":{"street":"peddar road","city":"Mumbai","state":"Maharashtra"}}`

	var modelData map[string]interface{}

	err = json.Unmarshal([]byte(modelJson), &modelData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(modelData)
	fmt.Println(modelData["model"])
	fmt.Println(modelData["age"])
	fmt.Println(modelData["address"])

	address := modelData["address"].(map[string]interface{})
	fmt.Println(address["street"])
	fmt.Println(address["city"])
	fmt.Println(address["state"])
}
