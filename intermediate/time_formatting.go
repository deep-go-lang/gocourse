package intermediate

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05Z07:00"
	date := "2025-04-05 09:16:40Z"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parsedTime)

	date2 := "April 5, 2025 9:22 AM"
	layout2 := "January 2, 2006 3:04 PM"

	parsedTime2, err := time.Parse(layout2, date2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parsedTime2)
	
	
	
	
	
}
