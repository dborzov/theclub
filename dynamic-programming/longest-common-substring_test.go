package main

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	P := "what a beautiful morning, there is something in the air"
	Q := "I love the smell of napalm in the morning, it is just the woman's things you wont get it"
	output := LCS(P, Q)
	fmt.Printf("lo and behold: %#v \n", output)
}
