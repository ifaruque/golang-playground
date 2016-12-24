package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t)
	fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()

	fmt.Println(t)
	fmt.Println(time.Now())

	week := 60 * 60 * 24 * 7 * 1e9 // must be in nanosecond
	duration := time.Duration(week)
	week_from_now_using_add := t.Add(duration)
	fmt.Println("Week from now using t.Add()", week_from_now_using_add)

	numberOfWeek := 2
	week_from_now_using_addDate := t.AddDate(0, 0, 7 * numberOfWeek)
	fmt.Println("Week from now using t.AddDate()", week_from_now_using_addDate)

	// formatting times:
	fmt.Println("RFC822:", t.Format(time.RFC822))     // 21 Dec 11 0852 UTC
	fmt.Println("ANSIC:", t.Format(time.ANSIC))      // Wed Dec 21 08:56:34 2011
	fmt.Println("String Formatter", t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}
