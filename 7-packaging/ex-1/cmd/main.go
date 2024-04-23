package main

import	(
	"github.com/henrymoreirasilva/go-expert/7-packaging/math"
	"github.com/goolge/uuid"
)

func main() {
	println("*** Packaging ***")


	m := math.Math{A: 1, B: 5}
	println(m.Add())

	println(m.Add())
	println(uuid.New().String())
}