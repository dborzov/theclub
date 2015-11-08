package main

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	P := "what a beautiful supermorning: there is something in the air"
	Q := "I love the smell of napalm in the lesupermorningmorning, it is just the woman's things you wont get it"
	output := LCS(P, Q)
	fmt.Printf("MATCH: [%#v] (length: %v) \n", output, len(output))
}
