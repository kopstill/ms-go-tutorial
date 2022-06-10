package concurrency

import "testing"

func TestFoo(t *testing.T) {
	// for i := 0; i < 10; i++ {
	foo()
	// }
}

func TestBoo(t *testing.T) {
	// for i := 0; i < 100; i++ {
	boo()
	// }
}

func TestChannel(t *testing.T) {
	channel()
}

func TestMultiPlexing(t *testing.T) {
	multiplexing()
}

func TestFibonacci(t *testing.T) {
	fibonacci()
}
