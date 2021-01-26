package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	var r1 = 1
	var r2 = 1
	if r1 != r2 {
		t.Error("Error during testing")
	}