package basics


import "fmt"

import "maps"

func main() {

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["key1"] = 1
	myMap["key2"] = 2
	myMap["key3"] = 3
	myMap["Rashmika"] = 25

	fmt.Println(myMap)

	fmt.Println(myMap["Rashmika"])

	fmt.Println(myMap["key4"])

	myMap["Rashmika"] = 500

	fmt.Println(myMap)

	delete(myMap, "key2")

	fmt.Println(myMap)

	fmt.Println(myMap["Rashmika"])

	myMap["key1"] = 900

	myMap["key3"] = 1000

	myMap["Rashmika"] = 2000

	fmt.Println(myMap)

	// clear(myMap)

	// fmt.Println(myMap)

	value, unknownValue := myMap["Rashmika"]

	fmt.Println(value, unknownValue)

	_, unknownValue2 := myMap["key3"]

	fmt.Println("There is a value for key3:", unknownValue2)

	myMapNew := map[string]int{
		"neha": 40,
		"kriti": 50,
		"tapsee": 60,
	}

	myMapNew1 := map[string]int{
		"neha": 70,
		"kriti": 80,
		"tapsee": 90,
	}

	myMapNew2 := map[string]int{
		"neha": 40,
		"kriti": 50,
		"tapsee": 60,
	}


	fmt.Println(myMapNew)
	fmt.Println(myMapNew1)
	fmt.Println(myMapNew2)

	for key, value := range myMapNew {
		fmt.Println(key, value)
	}
	
	for key, value := range myMapNew1 {
		fmt.Println(key, value)
	}

	for key, value := range myMapNew2 {
		fmt.Println(key, value)
	}

	if maps.Equal(myMapNew, myMapNew1) {
		fmt.Println("Both maps are equal")
	} else {
		fmt.Println("Both maps are not equal")
	}

	if maps.Equal(myMapNew, myMapNew2) {
		fmt.Println("Both maps are equal")
	} else {
		fmt.Println("Both maps are not equal")
	}

	for _, value := range myMapNew2 {
		fmt.Println(value)
	}

	 _, ok := myMapNew2["kriti"]

	 if ok {
		fmt.Println("Key exists")
	 } else {
		fmt.Println("Key does not exist")
	 }

	//  myNewMap3 := make(map[string]string)

	//  if myNewMap3 == nil {
	// 	fmt.Println("Map is nil")
	//  } else {
	// 	fmt.Println("Map is not nil")
	//  }

	//  myNewMap4 := map[string]string{}

	//  fmt.Println(myNewMap4)

	var myNewMap5 map[string]string

	if myNewMap5 == nil {
		fmt.Println("Map is nil")
	} else {
		fmt.Println("Map is not nil")
	}
	
	trade := myNewMap5["kriti"]

	if trade == "" {
		fmt.Println("trade does not exist")
	} else {
		fmt.Println("trade exists")
	}

	// myNewMap5["kriti"] = "trade"

	// fmt.Println(myNewMap5)

	myNewMap6 := make(map[string]string)

	myNewMap6["kriti"] = "trade"

	fmt.Println(myNewMap6)

	var extraMap map[string]string = map[string]string{
		"Prajakta": "Mali",
		"Rashmika": "Kannada",
		"Kriti": "Hindi",
		"Tapsee": "Telugu",
		"Neha": "Marathi",
	}

	fmt.Println(extraMap)

	// var actress map[string]string

	// actress["Rashmika"] = "beautiful"

	// fmt.Println(actress)

	// fmt.Println(actress["Rashmika"])

	// actress["Rashmika"] = "cute"

	// fmt.Println(actress)

	actress := make(map[string]string)

	actress["Rashmika"] = "beautiful"

	fmt.Println(actress)

	var models map[string]string = map[string]string{
		"Rashmika": "beautiful",
		"Kriti": "cute",
		"Tapsee": "pretty",
		"Neha": "gorgeous",
	}

	fmt.Println(models)

	fmt.Println("Length of models map is", len(models))

	for key, value := range models {
		fmt.Println(key, value)
	}

	angels := make(map[string]map[string]string)

	fmt.Println(angels)

	angels["girls"] = map[string]string{
		"Rashmika": "beautiful",
		"Kriti": "cute",
		"Tapsee": "pretty",
		"Neha": "gorgeous",
	}

	angels["women"] = map[string]string{
		"Kriti": "cute",
		"Tapsee": "pretty",
		"Neha": "gorgeous",
	}

	fmt.Println(angels)

	fmt.Println(angels["girls"]["Neha"])

	fmt.Println(angels["women"]["Tapsee"])

	angels["girls"]["Neha"] = "sexy"

	fmt.Println(angels)

	angels["hotties"] = models

	fmt.Println(angels)
	


	
	
	
	
	
	




	
	
	




	

	
	


}
