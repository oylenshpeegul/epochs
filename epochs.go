package epochs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"
)

const SecondsPerDay = 24 * 60 * 60

// Epoch gets a Unix time of the given x after dividing by q and
// adding s.
func epoch(x, q, s *big.Int) time.Time {
	z := new(big.Int)
	m := new(big.Int)
	z.DivMod(x, q, m)
	z.Add(z, s)
	r := m.Mul(m, big.NewInt(1e9)).Div(m, q)
	return time.Unix(z.Int64(), r.Int64()).UTC()
}

// Chrome time is the number of microseconds since 1601-01-01, which
// is 11,644,473,600 seconds before the Unix epoch.
func Chrome(num int64) time.Time {
	return epoch(
		big.NewInt(num),
		big.NewInt(1e6),
		big.NewInt(-11644473600),
	)
}

// Cocoa time is the number of seconds since 2001-01-01, which
// is 978,307,200 seconds after the Unix epoch.
func Cocoa(num int64) time.Time {
	return epoch(
		big.NewInt(num),
		big.NewInt(1),
		big.NewInt(978307200),
	)
}

// GoogleCalendar seems to count 32-day months from the day before the
// Unix epoch.
func GoogleCalendar(num int64) time.Time {

	n := int(num)

	totalDays := n / SecondsPerDay
	seconds := n % SecondsPerDay

	// A "Google month" has 32 days!
	months := totalDays / 32
	days := totalDays % 32

	// The "Google epoch" is apparently off by a day.
	t := time.Unix(-SecondsPerDay, 0).UTC()

	// Add the days first...
	u := t.AddDate(0, 0, days)

	// ...then the months.
	v := u.AddDate(0, months, 0)

	// ...then the seconds.
	w := v.Add(time.Duration(seconds * 1e9))

	return w
}

// ICQ time is the number of days since 1899-12-30, which is
// 2,209,161,600 seconds before the Unix epoch. Days can have a
// fractional part.
func ICQ(days float64) time.Time {

	t := time.Unix(-2209161600, 0).UTC()

	intdays := int(days)

	// Want the fractional part of the day in nanoseconds.
	fracday := int64((days - float64(intdays)) * SecondsPerDay * 1e9)

	return t.AddDate(0, 0, intdays).Add(time.Duration(fracday))
}

// Java time is the number of milliseconds since the Unix epoch.
func Java(num int64) time.Time {
	return epoch(
		big.NewInt(num),
		big.NewInt(1000),
		big.NewInt(0),
	)
}

// Mozilla time (e.g., formhistory.sqlite) is the number of
// microseconds since the Unix epoch.
func Mozilla(num int64) time.Time {
	return epoch(
		big.NewInt(num),
		big.NewInt(1e6),
		big.NewInt(0),
	)
}

// OLE time is the number of days since 1899-12-30, which is
// 2,209,161,600 seconds before the Unix epoch. Days can have a
// fractional part and is given as a string of hex characters
// representing an IEEE 8-byte floating-point number.
func OLE(days string) time.Time {
	var d [8]byte
	var f float64

	n, err := fmt.Sscanf(
		days,
		"%02x%02x%02x%02x%02x%02x%02x%02x",
		&d[0], &d[1], &d[2], &d[3], &d[4], &d[5], &d[6], &d[7],
	)
	if err != nil {
		fmt.Println("fmt.Sscanf failed:", err)
	}
	if n != 8 {
		fmt.Println("fmt.Sscanf did not scan 8 items:", n)
	}

	buf := bytes.NewReader(d[:])
	if err := binary.Read(buf, binary.LittleEndian, &f); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	return ICQ(f)
}

// Symbian time is the number of microseconds since the year 0, which
// is 62,167,219,200 seconds before the Unix epoch.
func Symbian(num int64) time.Time {
	return epoch(
		big.NewInt(num),
		big.NewInt(1e6),
		big.NewInt(-62167219200),
	)
}

// Unix time is the number of seconds since 1970-01-01.
func Unix(num int64) time.Time {
	return time.Unix(num, 0).UTC()
}

// UUID version 1 time (RFC 4122) is the number of hectonanoseconds
// (100 ns) since 1582-10-15, which is 12,219,292,800 seconds before
// the Unix epoch.
func UUIDv1(num int64) time.Time {
	return epoch(
		big.NewInt(num),
		big.NewInt(1e7),
		big.NewInt(-12219292800),
	)
}

// Windows date time (e.g., .NET) is the number of hectonanoseconds
// (100 ns) since 0001-01-01, which is 62,135,596,800 seconds before
// the Unix epoch.
func WindowsDate(num int64) time.Time {
	return epoch(
		big.NewInt(num),
		big.NewInt(1e7),
		big.NewInt(-62135596800),
	)
}

// Windows file time (e.g., NTFS) is the number of hectonanoseconds
// (100 ns) since 1601-01-01, which is 11,644,473,600 seconds before
// the Unix epoch.
func WindowsFile(num int64) time.Time {
	return epoch(
		big.NewInt(num),
		big.NewInt(1e7),
		big.NewInt(-11644473600),
	)
}
