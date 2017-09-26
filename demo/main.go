package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	// create a program that finds de number of days between two dates
	// datediff 2015-07-10 2015-07-11 debe dar: 1

	var date1, date2 string = os.Args[1], os.Args[2]

	t1, err := time.Parse("2006-01-02", date1)
	if err != nil {
		log.Fatalln(err)
	}

	t2, errno := time.Parse("2006-01-02", date2)
	if errno != nil {
		log.Fatalln(errno)
	}

	duration := t1.Sub(t2)
	diff := math.Abs(duration.Hours()) / 24

	if diff == 1 {
		fmt.Printf("Difference between dates is %d day.\n", int(diff))
	} else {
		fmt.Printf("Difference between dates is %d days.\n", int(diff))
	}
}
