package epochs

import (
	"testing"
	"time"
)

var EpochsTests = []struct {
	f   func(int64) time.Time
	num int64
	exp time.Time
}{
	{
		APFS,
		1234567890000000000,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		Chrome,
		12879041490000000,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		Chrome,
		12912187816559001,
		time.Date(2010, time.March, 4, 14, 50, 16, 559001000, time.UTC),
	},
	{
		Cocoa,
		256260690,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		Cocoa,
		314238233,
		time.Date(2010, time.December, 17, 0, 23, 53, 0, time.UTC),
	},
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
	{
		Java,
		1234567890000,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		Java,
		1283002533751,
		time.Date(2010, time.August, 28, 13, 35, 33, 751000000, time.UTC),
	},
	{
		Mozilla,
		1234567890000000,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		Symbian,
		63401787090000000,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
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
	{
		UUIDv1,
		134538606900000000,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		UUIDv1,
		0x1dc7711a73088f5,
		time.Date(2007, time.October, 10, 9, 17, 41, 739749300, time.UTC),
	},
	{
		WindowsDate,
		633701646900000000,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		WindowsDate,
		634496538123456789,
		time.Date(2011, time.August, 22, 23, 50, 12, 345678900, time.UTC),
	},
	{
		WindowsFile,
		128790414900000000,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
	},
	{
		WindowsFile,
		0x1cabbaa00ca9000,
		time.Date(2010, time.March, 4, 14, 50, 16, 559001600, time.UTC),
	},
}

func TestEpochs(t *testing.T) {
	for _, tt := range EpochsTests {

		obs := tt.f(tt.num)
		if obs != tt.exp {
			t.Errorf("%q => %q, want %q", tt.num, obs, tt.exp)
		}
	}
}
