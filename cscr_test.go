package cscr

import "testing"

func TestVersionNumber(t *testing.T) {
	want := "0.0.1"
	if got := VersionNumber(); got != want {
		t.Errorf("version number got %v want %v", got, want)
	}
}
