package datatype

import "fmt"

func mapdemo() {
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)
	fmt.Println("Bob's age is", studentsAge["bob"])

	studentsAge = make(map[string]int)
	fmt.Println(studentsAge)
	studentsAge["john"] = 34
	studentsAge["bob"] = 33
	fmt.Println(studentsAge)
	fmt.Println("Bob's age is", studentsAge["bob"])
	fmt.Println("Christy's age is", studentsAge["christy"])
	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age is not found")
	}

	var emptymap = make(map[string]int)
	fmt.Println(emptymap)
	fmt.Println(emptymap == nil)
	emptymap["param1"] = 1
	emptymap["param2"] = 2
	emptymap["param3"] = 3
	fmt.Println(emptymap)
	delete(emptymap, "param2")
	fmt.Println(emptymap)
	delete(emptymap, "param100")
	fmt.Println(emptymap)
	for k, v := range emptymap {
		fmt.Printf("%v\t%v\n", k, v)
	}
	for v := range emptymap {
		fmt.Println("key:", v)
	}

	var nilmap map[string]int
	fmt.Println(nilmap)
	fmt.Println(nilmap == nil)

	// var studentsAge1 map[string]int
	// studentsAge1["john"] = 32
	// studentsAge1["bob"] = 31
	// fmt.Println(studentsAge1)
}
