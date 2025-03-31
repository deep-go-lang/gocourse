package intermediate

import "fmt"

import "unicode/utf8"

func main() {

	message := "Hello, World!\nHello, World!\nHello, World!\nHello, World!"

	message2 := "Hello, world!\tHello, world!\tHello, world!\tHello, world!"

	message3 := "Hello, world! \rGo!"



	rawString := `
	Hello, World!
	Hello, World!
	Hello, World!
	Hello, World!
	`

	fmt.Println(message)
	fmt.Println(rawString)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println("the length of the message is", len(rawString))
	fmt.Println("the first character in message is", rawString[0])

	message4 := "Hello, 世界!"
	fmt.Println("the length of the message is", len(message4))
	fmt.Println("the first character in message is", message4[0])

	greeting := "Hello"

	name := "Aishwarya"

	msg := greeting + " " + name
	fmt.Println(msg)

	// str1 := "Aishwarya"

	str2 := "kriti"

	str3 := "Kriti"

	fmt.Println(str3 < str2)

	for i, c := range str3 {
		fmt.Printf("the letter at index %d is %c\n", i, c)
	}

	for _, c := range str3 {
		fmt.Printf("%d\n", c)
		fmt.Printf("%c\n", c)
		fmt.Printf("%T\n", c)
		fmt.Printf("%v\n", c)
		fmt.Printf("%x\n", c)
	}

	fmt.Println("Rune length of str3 is", len([]rune(str3)))

	fmt.Println("Rune count", utf8.RuneCountInString(str3))

	greetingWithName := greeting + " " + name

	fmt.Println(greetingWithName)

	var ch rune = '世'

	var ch2 rune = '界'

	var ch3 rune = '👋'

	var ch4 rune = 'p'

	fmt.Println(ch)
	fmt.Println(ch2)
	fmt.Println(ch3)
	fmt.Println(ch4)

	fmt.Println(string(ch))
	fmt.Println(string(ch2))
	
	fmt.Println(string(ch3))

	fmt.Println(string(ch4))

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", ch2)
	fmt.Printf("%c\n", ch3)
	fmt.Printf("%c\n", ch4)

	fmt.Printf("%T\n", ch)
	fmt.Printf("%T\n", ch2)
	fmt.Printf("%T\n", ch3)
	fmt.Printf("%T\n", ch4)

	fmt.Printf("%T\n", string(ch))

	const india = "卸ヌ宛"

	fmt.Println(india)

	morning := "何せて"

	for _, c := range morning {
		fmt.Printf("%c\n", c)
	}

	for _, c := range morning {
		fmt.Printf("%d\n", c)
	}


	for _, c := range morning {
		fmt.Printf("%T\n", c)
	}

	for _, c := range morning {
		fmt.Println(string(c))
	}

	for _, c := range morning {
		fmt.Printf("%x\n", c)
	}

	for _, c := range morning {
		fmt.Printf("%v\n", c)
	}

	fmt.Println("Rune count", utf8.RuneCountInString(morning))

	smily := "😀"

	fmt.Println(smily)

	fmt.Println(len(smily))	

	for _, c := range smily {
		fmt.Printf("%c\n", c)
	}

	for _, c := range smily {
		fmt.Printf("%d\n", c)
	}	

	for _, c := range smily {
		fmt.Printf("%T\n", c)
	}	

	for _, c := range smily {
		fmt.Println(string(c))
	}

	for _, c := range smily {
		fmt.Printf("%x\n", c)
	}

	for _, c := range smily {
		fmt.Printf("%v\n", c)
	}

	var p []rune = []rune(smily)

	fmt.Println(p)

	fmt.Println(string(p))


	myString := "Indian"

	for _, c := range myString {
		fmt.Println(c)
	}

	for _, c := range myString {
		fmt.Printf("%c\n", c)
	}

	

	
	
	
	
	
	
	
	
	
	
	
	

	
	

}
