package errorlog

import (
	"fmt"
	"log"
)

func simpleLog() {
	log.Print("Hey, I'm a log!")
	log.Fatal("Hey, I'm an error log!")
	// defer fmt.Print("Can you see me?")
	fmt.Print("Can you see me?")
}

func panicLog() {
	log.Panic("Hey, I'm an error log!")
	fmt.Print("Can you see me?")
}

func prefixLog() {
	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a log!")
	log.Fatal("Hey, I'm an error log!")
}
