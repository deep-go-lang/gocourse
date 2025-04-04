package intermediate

import (
	"fmt"
	"time"
)	

func main() {
	now := time.Now()
	fmt.Println(now)

	specificTime := time.Date(2025, time.April, 4, 17, 41, 50, 50000, time.UTC)
	fmt.Println(specificTime)

	parsedTime, _ := time.Parse(time.RFC3339, "2025-04-04T17:41:50Z")
	fmt.Println(parsedTime)

	// Mon Jan 2 15:04:05 MST 2006

	parsedTime1, _ := time.Parse("Mon 1 06 15:04", "Fri 4 25 17:41")
	fmt.Println(parsedTime1)

	parsedTime2, _ := time.Parse("06-Jan-2", "25-Apr-4")
	fmt.Println(parsedTime2)

	parsedTime3, _ := time.Parse("06.Jan.02", "25.Apr.20")
	fmt.Println(parsedTime3)

	parsedTime4, _ := time.Parse("06.Jan.02.15.04.05", "25.Apr.08.17.41.50")
	fmt.Println(parsedTime4)

	t := time.Now()
	fmt.Println(t)

	t1 := t.Format("06-Jan-02.04:15:05:Monday")
	fmt.Println(t1)

	oneDayLater := t.Add(24 * time.Hour)
	fmt.Println(oneDayLater)

	oneDayLater1 := t.AddDate(0, 0, 1)
	fmt.Println(oneDayLater1)
	fmt.Println(oneDayLater.Weekday())

	fmt.Println("Rounded Time:", t.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Date(2025, time.April, 4, 20, 29, 50, 50000, time.UTC)
	fmt.Println("Time in UTC:",t)

	time_local := t.In(loc)

	fmt.Println("Time in Kolkata:",time_local)

	roundedTime := t.Round(time.Hour)
	fmt.Println("Rounded Time in UTC:",roundedTime)

	roundedTimeLocal := roundedTime.In(loc)
	fmt.Println("Rounded Time in Kolkata:",roundedTimeLocal)

	tNow := time.Now()
	fmt.Println("Current Time:",tNow)

	fmt.Println("Truncated Time:",tNow.Truncate(time.Hour))

	locSwiss, _ := time.LoadLocation("Europe/Zurich")
	tSwiss := tNow.In(locSwiss)
	fmt.Println("Time in Zurich:",tSwiss)

	t3 := time.Date(2025, time.April, 4, 17, 41, 50, 50000, time.UTC)

	t4 := time.Date(2025, time.April, 4, 21, 41, 50, 50000, time.UTC)

	diff := t4.Sub(t3)
	fmt.Println("Difference:",diff)

	fmt.Println("t4 after t3", t4.After(t3))

	
	

	
	
	
	
	

	
	
	
	
	
	
	
	
	
	
	
	
}
