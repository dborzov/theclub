package main

import (
	"fmt"
	"testing"
)

func TestLoadBranch(t *testing.T) {
	examples := []string{
		"4",
		"(5,3,3)",
	}
	for _, e := range examples {
		obj, err := loadBranch([]byte(e))
		fmt.Printf("%v into %v\n", e, obj)
		if err != nil {
			fmt.Printf("  FAILED AT: %v\n", err)
		}
	}

}
