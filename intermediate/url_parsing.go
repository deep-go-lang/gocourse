package intermediate

import (
	"fmt"
	"net/url"
)

func main() {

	rawUrl := "https://www.example.com:8080/path?query=param#fragment"


	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Parsed URL:", parsedUrl)
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Query:", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	fmt.Println("--------------------------------")

	rawUrl2 := "https://www.example.com/path?name=Urvashi&age=20"

	parsedUrl2, err := url.Parse(rawUrl2)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Parsed URL:", parsedUrl2)
	fmt.Println("Scheme:", parsedUrl2.Scheme)
	fmt.Println("Host:", parsedUrl2.Host)
	fmt.Println("Port:", parsedUrl2.Port())
	fmt.Println("Path:", parsedUrl2.Path)
	fmt.Println("Query:", parsedUrl2.RawQuery)
	fmt.Println("Fragment:", parsedUrl2.Fragment)

	fmt.Println("----------------------------------")

	queryParams := parsedUrl2.Query()

	fmt.Println(queryParams)

	fmt.Println("Name:", queryParams.Get("name"))

	fmt.Println("Age:", queryParams.Get("age"))

	fmt.Println("----------------------------------")

	baseUrl := &url.URL{
		Scheme: "https",
		Host: "www.example.com",
		Path: "/path",
	}

	query := baseUrl.Query()

	query.Set("name", "Urvashi")
	query.Set("age", "28")
	baseUrl.RawQuery = query.Encode()

	fmt.Println("Full URL:", baseUrl.String())

	fmt.Println("----------------------------------")

	values := url.Values{}
	values.Add("name", "Urvashi")
	values.Add("age", "28")
	values.Add("city", "Mumbai")
	values.Add("country", "India")

	encodedQuery := values.Encode()

	fmt.Println("Encoded Query:", encodedQuery)

	fmt.Println("----------------------------------")

	fullUrl := "https://www.example.com/search?" + encodedQuery

	fmt.Println("Full URL:", fullUrl)

	fmt.Println("----------------------------------")

	
	

	




	

}
