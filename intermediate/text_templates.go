package intermediate

import (
	"log"
	"os"
	"text/template"
	"bufio"
	"fmt"
	"strings"
)

func main() {

	// tmpl := template.New("example")

	// tmpl, err := tmpl.Parse("Hello, {{.Name}}! How are you?\n")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data := map[string]interface{}{
	// 	"Name": []string{"John", "Jane", "Jim"},
	// }

	// for _, name := range data["Name"].([]string) {
	// 	err = tmpl.Execute(os.Stdout, map[string]interface{}{
	// 		"Name": name,
	// 	})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// tmpl2 := template.New("example2")

	// tmpl2 = template.Must(tmpl2.Parse("Hi, {{.Name}}! I love you\n"))

	// data2 := map[string]interface{}{
	// 	"Name": "Urvashi",
	// }

	// err = tmpl2.Execute(os.Stdout, data2)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name: ")

	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"Welcome": "Welcome, {{.Name}}! We are glad you joined us\n",
		"Notification": "{{.Name}}! You have a new notification: {{.Notification}}\n",
		"Error": "Oops! Something went wrong: {{.Error}}\n",
	}


	parsedTemplates := make(map[string]*template.Template)

	for tmplName, tmplStr := range templates {
		tmpl := template.Must(template.New(tmplName).Parse(tmplStr))
		parsedTemplates[tmplName] = tmpl
	}

	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		choice, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		choice = strings.TrimSpace(choice)

		var data map[string]interface{}

		var tmpl *template.Template

		switch choice {
		case "1":
			data = map[string]interface{}{
				"Name": name,
			}
			tmpl = parsedTemplates["Welcome"]

		case "2":
			fmt.Println("Enter your notification: ")
			notification, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			notification = strings.TrimSpace(notification)

			data = map[string]interface{}{
				"Name": name,
				"Notification": notification,
			}
			tmpl = parsedTemplates["Notification"]
			
		case "3":
			fmt.Println("Enter your error message: ")
			error, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			error = strings.TrimSpace(error)

			data = map[string]interface{}{
				"Name": name,
				"Error": error,
			}
			tmpl = parsedTemplates["Error"]
			
		case "4":
			fmt.Println("Exiting...")
			return
			
		default:
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		if err := tmpl.Execute(os.Stdout, data); err != nil {
			log.Fatal(err)
		}


		
		
		
		
	}
	








	


}