package datatype

import "fmt"

func slice() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))

	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}

	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2
	// remove := 4
	if remove < len(letters) {
		fmt.Println("Before", letters, "Remove", letters[remove])
		letters = append(letters[:remove], letters[remove+1:]...)
		fmt.Println("After", letters)
	}

	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])
	fmt.Println(slice2)

	letters = []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice11 := letters[0:2]
	// slice22 := letters[1:4]
	// slice22 := make([]string, 3, 3)
	slice22 := make([]string, 3, 4)
	copy(slice22, letters[1:4])

	slice11[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice22)
}
