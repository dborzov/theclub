package main

import (
	"fmt"
	"strconv"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) String() string {
	if t == nil {
		return "x"
	}
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Value)
	}
	return fmt.Sprintf("T(%v,L: %s, R: %s)", t.Value, t.Left, t.Right)
}
