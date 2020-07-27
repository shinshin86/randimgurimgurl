package randimgurimgurl

import (
	"testing"
)

func TestGetter(t *testing.T) {
	imgurl := Getter()

	if imgurl == "" {
		t.Fatal("failed test")
	}
}
