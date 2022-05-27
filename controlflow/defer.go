package controlflow

import (
	"fmt"
	"os"
)

func deferDemo() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}
}

func deferDemo1() {
	newfile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}
	defer newfile.Close()

	if _, error := newfile.WriteString("Learning Go!"); error != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}
	newfile.Sync()
}
