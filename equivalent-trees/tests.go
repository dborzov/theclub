package main

import (
	"errors"
	"strconv"
)

// serialized trees in quasy lisp syntax
const (
	ex1 = `(3,(1,1,1),(8,5,13))`
	ex2 = `(8,(3,(1,1,2),5),13)`
)

func loadBranch(s []byte) (*Tree, error) {
	if len(s) == 0 {
		return nil, nil
	}

	if s[0] != byte('(') {
		val, err := strconv.Atoi(string(s))
		if err != nil {
			return nil, err
		}
		return &Tree{Value: val}, nil
	}

	lb := -1
	rb := -1
	// lb is the position on the first comma if exists
	for i := 1; i < len(s); i++ {
		if s[i] == byte(',') {
			lb = i
			break
		}
	}

	if lb == -1 {
		return nil, errors.New("Invalid syntax, no comma found")
	}

	for j := len(s); j > lb; j-- {
		if s[j] == byte(',') {
			rb = j
			break
		}
	}

	if rb == -1 {
		return nil, errors.New("Invalid syntax, no second comma found")
	}

	val, err := strconv.Atoi(string(s[1:lb]))
	if err != nil {
		return nil, errors.New("Cannot recognize the value")
	}

	left, err := loadBranch(s[lb:rb])
	if err != nil {
		return nil, err
	}

	right, err := loadBranch(s[rb : len(s)-1])
	if err != nil {
		return nil, err
	}

	return &Tree{
		Value: val,
		Left:  left,
		Right: right,
	}, nil
}
