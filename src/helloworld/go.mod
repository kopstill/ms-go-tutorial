module helloworld

go 1.18

require github.com/kopever/calculator v0.0.0

replace github.com/kopever/calculator => ../calculator

require kopstill/calculator v1.1.1

replace kopstill/calculator => ../package/calculator
