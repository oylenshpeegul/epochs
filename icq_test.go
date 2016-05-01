package epochs

import (
	"testing"
	"time"
)

var ICQTests = []struct {
	f   func(float64) time.Time
	num float64
	exp time.Time
}{
	{
		ICQ,
		0,
		time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC),
	},
	{
		ICQ,
		41000,
		time.Date(2012, time.April, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ICQ,
		41056.2752083333,
		time.Date(2012, time.May, 27, 6, 36, 17, 999997418, time.UTC),
	},
	{
		ICQ,
		41056.2967361111,
		time.Date(2012, time.May, 27, 7, 7, 17, 999999080, time.UTC),
	},
}

func TestICQ(t *testing.T) {
	for _, tt := range ICQTests {

		obs := tt.f(tt.num)
		if obs != tt.exp {
			t.Errorf("%q(%q) => %q, want %q", tt.f, tt.num, obs, tt.exp)
		}
	}
}
