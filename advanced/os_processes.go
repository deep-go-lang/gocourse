package advanced

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("ls", "-l")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(out))

	// pr, pw := io.Pipe()

	// cmd := exec.Command("grep", "hello")

	// cmd.Stdin = pr

	// go func() {
	// 	defer pw.Close()
	// 	pw.Write([]byte("hello world\nhello there\nUrvashi Rautela\n"))
	// }()

	// out, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// fmt.Println(string(out))


	// cmd := exec.Command("printenv", "SHELL")

	// out, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// fmt.Println(string(out))

	// cmd := exec.Command("sleep", "60")

	// err := cmd.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// time.Sleep(time.Second * 2)

	// err = cmd.Process.Kill()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// err = cmd.Wait()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// fmt.Println("Command Killed")

	// cmd := exec.Command("grep", "hello")

	// cmd.Stdin = strings.NewReader("hello world\nhello there\nUrvashi Rautela\n")


	// out, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// fmt.Println(string(out))

	// cmd := exec.Command("echo", "Hello, World!")

	// out, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(out))
}
