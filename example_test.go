package npmdl_test

import (
	"fmt"

	"github.com/bojand/go-npmdl"
)

// Example using the Client and getting Point download counts for all of npm for last day
func Example_Point() {
	out, _ := npmdl.GetPointCounts("", "")
	fmt.Printf("%v\n", out)
}

// Example using the Client and getting Point download counts for all of npm for last week
func Example_Point_Last_Week() {
	out, _ := npmdl.GetPointCounts("lask-week", "")
	fmt.Printf("%v\n", out)
}

// Example using the Client and getting Point download counts for express for a specific range
func Example_Point_Last_Week_Express() {
	out, _ := npmdl.GetPointCounts("2017-08-10:2017-08-20", "express")
	fmt.Printf("%v\n", out)
}

// Example using the Client and getting Range download counts for express for last week
func Example_Range_Last_Week_Express() {
	out, _ := npmdl.GetRangeCounts("last-week", "express")
	fmt.Printf("%v\n", out)
}
