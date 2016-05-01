package epochs

import (
	"testing"
	"time"
)

var OLETests = []struct {
	f   func(string) time.Time
	num string
	exp time.Time
}{
	{
		OLE,
		"dedddd5d3f76e340",
		time.Date(2009, time.February, 13, 23, 31, 30, 83, time.UTC),
	},
	{
		OLE,
		"8ad371b4bcd2e340",
		time.Date(2011, time.February, 23, 21, 31, 43, 127000061, time.UTC),
	},
}

func TestOLE(t *testing.T) {
	for _, tt := range OLETests {
		obs := tt.f(tt.num)
		if obs != tt.exp {
			t.Errorf("%q(%q) => %q, want %q", tt.f, tt.num, obs, tt.exp)
		}
	}
}
