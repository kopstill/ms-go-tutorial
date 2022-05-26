package controlflow

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	ifDemo()
}

func TestSwitch(t *testing.T) {
	switchDemo()
}

func TestLocation(t *testing.T) {
	region, continent := location("Delhi")
	fmt.Printf("John works in %s, %s\n", region, continent)
}

func TestLearn(t *testing.T) {
	learn()
}

func TestMatch(t *testing.T) {
	match()
}

func TestFooTime(t *testing.T) {
	fooTime()
}

func TestCompare(t *testing.T) {
	compare()
}
