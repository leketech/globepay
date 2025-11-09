package unit

import (
	"testing"
)

func TestSimple(t *testing.T) {
	// Simple test to verify the testing framework works
	if 1+1 != 2 {
		t.Error("Basic math failed")
	}
}