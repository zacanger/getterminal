package getterminal

import (
	"os"
	"testing"
)

func TestGetTerminal(t *testing.T) {
	exp := "st"
	os.Setenv("TERMINAL", exp)
	if got := GetTerminal(); got != exp {
		t.Errorf("Expected %q, got %q", exp, got)
	}

	exp = "x-terminal-emulator"
	os.Setenv("TERMINAL", "asdfasdf")
	if got := GetTerminal(); got != exp {
		t.Errorf("Expected %q, got %q", exp, got)
	}

}
