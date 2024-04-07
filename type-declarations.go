package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var NoKTPuj NoKTP = "16494616431305406"
	var MarriedStatus = true

	fmt.Println(NoKTPuj)
	fmt.Println(MarriedStatus)
}
