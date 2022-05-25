module helloworld

go 1.18

require github.com/kopever/calculator v0.0.0

replace github.com/kopever/calculator => ../calculator

require (
	kopstill/calculator v1.1.1
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace kopstill/calculator => ../package/calculator
