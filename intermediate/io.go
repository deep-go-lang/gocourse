package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
}


func cloeseResource(r io.Closer) {
	err := r.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func bufferExample() {
	var buf bytes.Buffer
	buf.WriteString("Hello, World!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	reader1 := strings.NewReader("Hello, ")
	reader2 := strings.NewReader("World!")
	multiReader := io.MultiReader(reader1, reader2)
	buf := new(bytes.Buffer)
	n, err := 	buf.ReadFrom(multiReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	fmt.Println(buf.String())
}



func main() {

	fmt.Println("------Read From Reader------")
	readFromReader(strings.NewReader("Hello, World!"))

	fmt.Println("------Write To Writer------")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello, World!")
	fmt.Println(writer.String())

	fmt.Println("------Buffer Example------")
	bufferExample()

	fmt.Println("------Multi Reader Example------")
	multiReaderExample()

	// fmt.Println("------Close Resource Example------")
	// cloeseResource(os.Stdout)

	fmt.Println("------Pipe Example------")
	pipeExample()

	fmt.Println("------Write To File Example------")
	writeToFile("test.txt", "Hello, Urvashi!")

	fmt.Println("------Copy Example------")
	copyExample()

}

func pipeExample() {
	reader, writer := io.Pipe()
	go func() {
		writer.Write([]byte("Hello, World!"))
		writer.Close()
	}()
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	fmt.Println(buf.String())
	
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer cloeseResource(file)
	n, err := file.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	// writer := io.Writer(file)
	// n, err =  writer.Write([]byte("Hello, World!"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(n)
	
}

func copyExample() {
	reader := strings.NewReader("Hello, Aishwarya!\n")
	writer := os.Stdout
	n, err := io.Copy(writer, reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

