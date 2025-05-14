package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	password := "Urvashi@123"

	// hash := sha256.Sum256([]byte(password))

	// hash1 := sha512.Sum512([]byte(password))

	// fmt.Println(hash)

	// fmt.Println(hash1)

	// fmt.Println(password)

	// fmt.Printf("SHA256 Hash hex: %x\n", hash)

	// fmt.Printf("SHA256 Hash base64: %s\n", base64.StdEncoding.EncodeToString(hash[:]))

	// fmt.Printf("SHA512 Hash hex: %x\n", hash1)

	// fmt.Printf("SHA512 Hash base64: %s\n", base64.StdEncoding.EncodeToString(hash1[:]))

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Salt:", salt)

	fmt.Printf("Original Salt: %x\n", salt)

	signupSaltHashedPassword := hashPassword(password, salt)
	fmt.Println("Hashed Password:", signupSaltHashedPassword)

	saltBase64 := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt Base64:", saltBase64)


	deoodeSalt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Deode Salt:", deoodeSalt)

	loginHash := hashPassword("Urvashi@123", deoodeSalt)

	fmt.Println("Login Password:", loginHash)

	if signupSaltHashedPassword == loginHash {
		fmt.Println("Passwrod is Correct")
	} else {
		fmt.Println("Passwrod is Incorrect")
	}

	fmt.Println("------------------------------")

	hashOfOnlyPassword := sha256.Sum256([]byte(password))

	fmt.Println(hashOfOnlyPassword)

	fmt.Println("String of hash of only password without salt:", base64.StdEncoding.EncodeToString(hashOfOnlyPassword[:]))


}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}


