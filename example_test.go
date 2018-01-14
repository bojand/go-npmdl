package npmdl_test

import (
	"fmt"

	"github.com/bojand/go-npmdl"
)

// Example using GetPointCounts to get point download counts for express for a specific range
func Example() {
	out, _ := npmdl.GetPointCounts("2017-08-10:2017-08-20", "express")
	fmt.Printf("%+v", out)
}

// Example using GetPointCounts to get point download counts for all packages for last week
func Example_period() {
	out, _ := npmdl.GetPointCounts("lask-week", "")
	fmt.Printf("%+v", out)
}
