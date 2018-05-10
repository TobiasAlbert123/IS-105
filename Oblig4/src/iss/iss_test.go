package main

import (
	"testing"
	"fmt"
)

func TestRunAndOpen(t *testing.T) {
	keysUsed, keys := RunForTest()
	fmt.Println(keysUsed, len(keys))
	if outOfAPIKeys(keysUsed, keys) {
		t.Fail()
	}
	if invalidData {
		t.Fail()
	}
}

func outOfAPIKeys(keysUsed int, keys []string) bool{
	if keysUsed >= len(keys) {
		return true
	}

	return false
}