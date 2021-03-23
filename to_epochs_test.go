package epochs

import (
	"testing"
	"time"
)

var ToEpochsTests = []struct {
	f   func(time.Time) int64
	t   time.Time
	exp int64
}{
	{
		ToAPFS,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
		1234567890000000000,
	},
	{
		ToChrome,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
		12879041490000000,
	},
	{
		ToChrome,
		time.Date(2010, time.March, 4, 14, 50, 16, 559001000, time.UTC),
		12912187816559001,
	},
	{
		ToCocoa,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
		256260690,
	},
	{
		ToCocoa,
		time.Date(2010, time.December, 17, 0, 23, 53, 0, time.UTC),
		314238233,
	},
	{
		ToJava,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
		1234567890000,
	},
	{
		ToJava,
		time.Date(2010, time.August, 28, 13, 35, 33, 751000000, time.UTC),
		1283002533751,
	},
	{
		ToMozilla,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
		1234567890000000,
	},
	{
		ToSymbian,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
		63401787090000000,
	},
	{
		ToUnix,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
		1234567890,
	},
	{
		ToUnix,
		time.Date(1930, time.November, 18, 0, 28, 30, 0, time.UTC),
		-1234567890,
	},
	{
		ToUUIDv1,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
		134538606900000000,
	},
	{
		ToUUIDv1,
		time.Date(2007, time.October, 10, 9, 17, 41, 739749300, time.UTC),
		0x1dc7711a73088f5,
	},
	{
		ToWindowsDate,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
		633701646900000000,
	},
	{
		ToWindowsDate,
		time.Date(2011, time.August, 22, 23, 50, 12, 345678900, time.UTC),
		634496538123456789,
	},
	{
		ToWindowsFile,
		time.Date(2009, time.February, 13, 23, 31, 30, 0, time.UTC),
		128790414900000000,
	},
	{
		ToWindowsFile,
		time.Date(2010, time.March, 4, 14, 50, 16, 559001600, time.UTC),
		0x1cabbaa00ca9000,
	},

	// This conversion to NTFS time will not preserve nanosecond
	// accuracy, but that's normal. It shouldn't complain about it.
	{
		ToWindowsFile,
		time.Date(2020, 3, 23, 10, 17, 0, 123456750, time.UTC),
		132294322201234567,
	},
}

func TestToEpochs(t *testing.T) {
	for _, tt := range ToEpochsTests {

		obs := tt.f(tt.t)
		if obs != tt.exp {
			t.Errorf("%q => %q, want %q", tt.t, obs, tt.exp)
		}
	}
}
