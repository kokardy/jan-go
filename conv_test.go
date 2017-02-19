package main

import (
	"testing"
)

var gs1codes []string = []string{
	"0114987058690557",
	"0114987058690038",
	"0114987058690113",
	"0114987058691554",
	"0114987058691035",
}

var jancodes []string = []string{
	"4987058690550",
	"4987058690031",
	"4987058690116",
	"4987058691557",
	"4987058691038",
}

func TestConvert(t *testing.T) {
	for i, gs1code := range gs1codes {
		gs1 := GS1(gs1code)
		jan := JAN(jancodes[i])
		result := gs1.ToJAN()
		if result != jan {
			t.Fatalf(
				"Converting GS1 to JAN failed: %s is expected. but %s",
				jan,
				result,
			)
		}
	}
}
