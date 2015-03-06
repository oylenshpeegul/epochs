package epochs

import (
	"testing"
	"time"
)

var ChromeTests = []struct {
	num int64
	exp time.Time
}{
	{12879041490000000, time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC)},
}

func TestChrome(t *testing.T) {

	for _, tt := range ChromeTests {

		obs := Chrome(tt.num)
		if obs != tt.exp {
			t.Errorf("Unix(%q) => %q, want %q", tt.num, obs, tt.exp)
		}
	}
}

var UnixTests = []struct {
	num int64
	exp time.Time
}{
	{1234567890, time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC)},
}

func TestUnix(t *testing.T) {
	for _, tt := range UnixTests {

		obs := Unix(tt.num)
		if obs != tt.exp {
			t.Errorf("Unix(%q) => %q, want %q", tt.num, obs, tt.exp)
		}
	}
}
