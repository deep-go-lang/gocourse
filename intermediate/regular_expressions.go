package intermediate

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("She said: \"Hello, World!\"")

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	email1 := "test@example.com"
	email2 := "invalid-email"

	match1 := re.MatchString(email1)
	match2 := re.MatchString(email2)

	fmt.Println(match1)
	fmt.Println(match2)

	re2 := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	date := "2024-01-01"

	match := re2.FindStringSubmatch(date)

	fmt.Println(match)

	source_string := "Hello, Urvashi! This is a test."

	re3 := regexp.MustCompile(`[aeiou]`)

	replaced_string := re3.ReplaceAllString(source_string, "-")

	fmt.Println(replaced_string)

	re4 := regexp.MustCompile(`(?i)Go`)

	text := "god is going to do good"

	matchNew := re4.MatchString(text)

	fmt.Println(matchNew)

	
	
	
	
	
	
	

}
