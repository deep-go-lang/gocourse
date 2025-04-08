package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("H~llo, Urvashi!\nYou ar~ a b~autiful girl!\n")

	encoded := base64.StdEncoding.EncodeToString(data)

	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(decoded))

	fmt.Println("--------------------------------")

	encoded = base64.URLEncoding.EncodeToString(data)

	fmt.Println(encoded)

	decoded, err = base64.URLEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(decoded))


}
