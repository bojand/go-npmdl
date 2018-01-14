package npmdl_test

import (
	"fmt"

	"github.com/bojand/go-npmdl"
)

// Example using GetPointCounts to get point download counts for all packages for last day
func ExampleGetPointCounts_all_counts_default_last_day() {
	out, _ := npmdl.GetPointCounts("", "")
	fmt.Printf("%v\n", out)
}

// Example using GetPointCounts to get point download counts for all packages for last week
func ExampleGetPointCounts_last_week() {
	out, _ := npmdl.GetPointCounts("lask-week", "")
	fmt.Printf("%v\n", out)
}

// Example using GetPointCountsto get point download counts for express for a specific range
func ExampleGetPointCounts_last_week_express() {
	out, _ := npmdl.GetPointCounts("2017-08-10:2017-08-20", "express")
	fmt.Printf("%v\n", out)
}

// Example using GetRangeCounts get point download counts for express for last week
func ExampleGetRangeCounts_last_week_express() {
	out, _ := npmdl.GetRangeCounts("last-week", "express")
	fmt.Printf("%v\n", out)
}
