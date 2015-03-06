package epochs

import (
	"testing"
	"time"
)

var MixedTests = []struct {
	f   func(int64) time.Time
	num int64
	exp time.Time
}{
	{
		Chrome,
		12879041490000000,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		Chrome,
		12912187816559001,
		time.Date(2010, time.March, 4, 14, 50, 16, 559001000, time.UTC)},
	{
		Unix,
		1234567890,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		Unix,
		-1234567890,
		time.Date(1930, time.November, 18, 0, 28, 30, 0, time.UTC),
	},
}

func TestMixed(t *testing.T) {
	for _, tt := range MixedTests {

		obs := tt.f(tt.num)
		if obs != tt.exp {
			t.Errorf("%q(%q) => %q, want %q", tt.f, tt.num, obs, tt.exp)
		}
	}
}
