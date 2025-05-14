package intermediate

import (
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// err := os.Mkdir("mydir1", 0755)
	// checkError(err)

	// err = os.Mkdir("mydir", 0755)
	// checkError(err)


	// data := []byte("")

	// os.WriteFile("mydir1/file", data, 0755)

	// err = os.MkdirAll("mydir/parent/child", 0755)
	// checkError(err)

	// err = os.MkdirAll("mydir/parent/child1", 0755)
	// checkError(err)

	// err = os.MkdirAll("mydir/parent/child2", 0755)
	// checkError(err)

	// err = os.MkdirAll("mydir/parent/child3", 0755)
	// checkError(err)

	// os.WriteFile("mydir/parent/file", data, 0755)

	// os.WriteFile("mydir/parent/child/file", data, 0755)

	// defer os.RemoveAll("mydir")

	// files, err := os.ReadDir("mydir/parent")
	// checkError(err)

	// for _, file := range files {
	// 	fmt.Println(file.Name(), file.IsDir(), file.Type())
	// }

	// fmt.Println("--------------------------------")

	// err = os.Chdir("mydir/parent/child")
	// checkError(err)

	// wd, err := os.Getwd()
	// checkError(err)

	// fmt.Println(wd)

	// files, err = os.ReadDir(".")
	// checkError(err)

	// for _, file := range files {
	// 	fmt.Println(file.Name(), file.IsDir(), file.Type())
	// }


	// fmt.Println("--------------------------------")

	// err = os.Chdir("../../..")
	// checkError(err)

	// wd, err = os.Getwd()
	// checkError(err)

	// fmt.Println(wd)

	// files, err = os.ReadDir(".")
	// checkError(err)

	// for _, file := range files {
	// 	fmt.Println(file.Name(), file.IsDir(), file.Type())
	// }

	// fmt.Println("--------------------------------")

	// pathfile := "mydir"

	// err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println(path, d.Name(), d.IsDir(), d.Type())
	// 	info, err := d.Info()
	// 	checkError(err)
	// 	fmt.Println(info.Size())
	// 	return nil
	// })
	// checkError(err)

	err := 	os.Remove("mydir1")
	checkError(err)



	

	
	
	
	
	

}
