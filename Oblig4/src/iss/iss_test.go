package main

import (
	"testing"
)

func TestRunAndOpen(t *testing.T) {
	runForTest()
	if globalError != "" {
		t.Errorf(globalError)
	}
	if outOfAPIKeys(GeoKeysUsed, len(SliceOfGeoKeys)) {
		t.Fail()
	}
	//this var is set to true during 'formatJson()' if something is wrong with ISS API data
	if invalidData {
		t.Fail()
	}
}

//checks if all API keys have been used, rendering the application useless
func outOfAPIKeys(used, available int) bool{
	if used >= available {
		return true
	}
	return false
}