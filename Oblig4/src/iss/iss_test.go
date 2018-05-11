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
		t.Error("Out of api keys\n")
	}
	//this var is set to true during 'formatJson()' if something is wrong with ISS API data
	if invalidData {
		t.Error("The ISS API data was invalid")
	}
}

//checks if all API keys have been used, rendering the application useless
func outOfAPIKeys(used, available int) bool{
	if used >= available {
		return true
	}
	return false
}