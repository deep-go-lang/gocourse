package intermediate

import (
	"fmt"
// 	"strconv"
	"strings"
// 	"regexp"
// 	"unicode/utf8"
)


func main() {

	// str1 := "Hello, Urvashi!"

	// str2 := "Hello, Aishwarya!"

	// str3 := str1 + " " + str2

	// fmt.Println(str3)

	// fmt.Println(len(str1))

	// fmt.Println(len(str2))

	// fmt.Println(len(str3))

	// fmt.Println(str3[0])

	// fmt.Println(str3[1:8])

	// num := 123

	// str4 := strconv.Itoa(num)

	// fmt.Println(len(str4))

	// actresses := "Kareena Kapoor-Deepika Padukone-Katrina Kaif-Priyanka Chopra"

	// actress := strings.Split(actresses, "-")

	// fmt.Println(actress)

	// fmt.Println(actress[0])

	// fmt.Println(actress[1])

	// fmt.Println(actress[2])

	// fmt.Println(actress[3])

	// cities := []string{"Mumbai", "Delhi", "Bangalore", "Hyderabad", "Chennai"}

	// joinedCities := strings.Join(cities, ", ")

	// fmt.Println(joinedCities)

	// fmt.Println(strings.Contains(joinedCities, "Mumbai"))

	// fmt.Println(strings.Replace(joinedCities, "Mumbai", "Pune", 1))

	// fmt.Println(joinedCities)

	// urvashi := "   Urvashi"

	// fmt.Println(urvashi)

	// fmt.Println(strings.TrimSpace(urvashi))

	// fmt.Println(strings.ToLower(actresses))

	// fmt.Println(strings.ToUpper(actresses))

	// fmt.Println(strings.Repeat(actresses, 3))

	// fmt.Println(strings.Count(actresses, "pr"))

	// fmt.Println(strings.HasPrefix(actresses, "Kar"))

	// fmt.Println(strings.HasSuffix(actresses, "Chopra"))

	// str5 := "Hello 777, Urvashi 555444!"

	// re := regexp.MustCompile(`\d+`)

	// fmt.Println(re.FindAllString(str5, 1))

	// str6 := "愛してるよウルヴァシ"

	// fmt.Println(utf8.RuneCountInString(str6))

	// fmt.Println(strings.ReplaceAll(joinedCities, "Mumbai", "Pune"))

	// fmt.Println(strings.Count(joinedCities, "Mumbai"))

	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")

	result := builder.String()

	fmt.Println(result)

	builder.WriteRune('!')

	builder.WriteRune(' ')

	builder.WriteString("Urvashi")

	result = builder.String()

	fmt.Println(result)

	builder.Reset()

	builder.WriteString("Starting Fresh!")

	result = builder.String()

	fmt.Println(result)



	

	
	
	
	
	
	
	
	
	
	



	

	
	
	
}
