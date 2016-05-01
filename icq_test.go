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
			t.Errorf("%f => %q, want %q", tt.num, obs, tt.exp)
		}
	}
}

var ToICQTests = []struct {
	f   func(time.Time) float64
	t   time.Time
	exp float64
}{
	{
		ToICQ,
		time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC),
		0,
	},
	{
		ToICQ,
		time.Date(2012, time.April, 1, 0, 0, 0, 0, time.UTC),
		41000,
	},
	{
		ToICQ,
		time.Date(2012, time.May, 27, 6, 36, 17, 999997418, time.UTC),
		41056.2752083333,
	},
	{
		ToICQ,
		time.Date(2012, time.May, 27, 7, 7, 17, 999999080, time.UTC),
		41056.2967361111,
	},
}

func TestToICQ(t *testing.T) {
	for _, tt := range ToICQTests {

		obs := tt.f(tt.t)
		if obs != tt.exp {
			t.Errorf("%q => %f, want %f", tt.t, obs, tt.exp)
		}
	}
}
