package epochs

import (
	"testing"
	"time"
)

var ChromeTests = []struct {
	num int64
	exp time.Time
}{
	{12879041490000000,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC)},
	{12912187816559001,
		time.Date(2010, time.March, 4, 14, 50, 16, 559001000, time.UTC)},
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
	{1234567890,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC)},
	{-1234567890,
		time.Date(1930, time.November, 18, 0, 28, 30, 0, time.UTC)},
}

func TestUnix(t *testing.T) {
	for _, tt := range UnixTests {

		obs := Unix(tt.num)
		if obs != tt.exp {
			t.Errorf("Unix(%q) => %q, want %q", tt.num, obs, tt.exp)
		}
	}
}
