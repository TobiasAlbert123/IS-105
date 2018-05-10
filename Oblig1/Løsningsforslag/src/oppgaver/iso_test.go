package oppgaver

import (
	"testing"
)

func TestGreetingExtendedASCII(t *testing.T) {
	ascii := ExtendedASCIIText()
	if !(isExtendedASCII(ascii)){
		t.Fail()
	}
}

func isExtendedASCII(s string) bool {
	for _, c := range s {
		if c > 255 {
			return false
		}
		if c < 127 {
			return false
		}
	}
	return true
}
