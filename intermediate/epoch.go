package intermediate

import (
	"fmt"
	"time"
)

func main() {
	epoch := time.Now().Unix()
	fmt.Println(epoch)

	time := time.Unix(epoch, 0)
	fmt.Println(time)

	fmt.Println(time.Format("2006-01-02 15:04:05"))

	
	
}
