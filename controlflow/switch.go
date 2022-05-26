package controlflow

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func switchDemo() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)
	fmt.Println(i)

	switch i {
	case 0:
		fmt.Println("zero...")
	case 1:
		fmt.Println("one...")
	case 2:
		fmt.Println("two...")
	default:
		fmt.Println("not match...")
	}

	fmt.Println("ok")
}

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func learn() {
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go!")
	default:
		fmt.Println("It's weekend, time to rest!")
	}
	fmt.Println(time.Now().Weekday())
}

func match() {
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	// contact := "9273987654"
	contact := "foo@bar.com"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}
}

func fooTime() {
	rand.Seed(time.Now().Unix())
	r := rand.Float64()
	fmt.Println(r)
	switch {
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}
}

func compare() {
	i := 15
	switch {
	case i > 100:
		fmt.Println("i is greater than 100")
		fallthrough
	case i > 10:
		fmt.Println("i is greater than 10")
		fallthrough
	case i < 0:
		fmt.Println("i is less than 0")
		fallthrough
	case i < 100:
		fmt.Println("i is less than 100")
	default:
		fmt.Println("compare default")
	}
}
