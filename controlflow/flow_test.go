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

func TestFor(t *testing.T) {
	forDemo()
}

func TestFor1(t *testing.T) {
	forDemo1()
}

func TestFor2(t *testing.T) {
	forDemo2()
}

func TestFor3(t *testing.T) {
	forDemo3()
}

func TestDefer(t *testing.T) {
	deferDemo()
}

func TestDefer1(t *testing.T) {
	deferDemo1()
}
