package main

import (
	"fmt"
	"testing"
)

func TestLoadBranch(t *testing.T) {
	examples := []string{
		"4",
		"(5,3,3)",
		`(3,(1,0,2),(8,5,13))`,
		`(8,(3,(1,1,2),5),13)`,
	}
	for _, e := range examples {
		obj, err := loadBranch([]byte(e))
		fmt.Printf("  --->  %v into %s\n", e, obj)
		if err != nil {
			fmt.Printf("  FAILED AT: %v\n", err)
		}
	}

}
