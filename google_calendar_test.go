package epochs

import (
	"testing"
	"time"
)

var GoogleCalendarTests = []struct {
	f   func(int64) time.Time
	num int64
	exp time.Time
}{
	{
		GoogleCalendar,
		1297899090,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		GoogleCalendar,
		1234567890,
		time.Date(2007, time.March, 16, 23, 31, 30, 0, time.UTC),
	},
}

func TestGoogleCalendar(t *testing.T) {
	for _, tt := range GoogleCalendarTests {

		obs := tt.f(tt.num)
		if obs != tt.exp {
			t.Errorf("%q(%q) => %q, want %q", tt.f, tt.num, obs, tt.exp)
		}
	}
}
