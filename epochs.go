package epochs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"time"
)

const secondsPerDay = 24 * 60 * 60
const nanosecondsPerDay = secondsPerDay * 1e9

// epoch2time gets a Unix time of the given x after dividing by q and
// adding s.
func epoch2time(x, q, s *big.Int) time.Time {
	z := new(big.Int)
	m := new(big.Int)
	z.DivMod(x, q, m)
	z.Add(z, s)
	r := m.Mul(m, big.NewInt(1e9)).Div(m, q)
	return time.Unix(z.Int64(), r.Int64()).UTC()
}

// time2epoch reverses epoch2time.
func time2epoch(t time.Time, m, s *big.Int) int64 {
	bf := new(big.Float).SetInt(big.NewInt(t.UnixNano()))
	bf.Quo(bf, big.NewFloat(1e9))
	bf.Sub(bf, new(big.Float).SetInt(s))
	bf.Mul(bf, new(big.Float).SetInt(m))

	r, acc := bf.Int64()
	if acc != big.Exact {
		fmt.Println(acc)
	}

	return r
}

// Chrome time is the number of microseconds since 1601-01-01, which
// is 11,644,473,600 seconds before the Unix epoch.
func Chrome(num int64) time.Time {
	return epoch2time(
		big.NewInt(num),
		big.NewInt(1e6),
		big.NewInt(-11644473600),
	)
}

// ToChrome returns the Chrome time for the given time.Time.
func ToChrome(t time.Time) int64 {
	return time2epoch(
		t,
		big.NewInt(1e6),
		big.NewInt(-11644473600),
	)
}

// Cocoa time is the number of seconds since 2001-01-01, which
// is 978,307,200 seconds after the Unix epoch.
func Cocoa(num int64) time.Time {
	return epoch2time(
		big.NewInt(num),
		big.NewInt(1),
		big.NewInt(978307200),
	)
}

// ToCocoa returns the Cocoa time for the given time.Time.
func ToCocoa(t time.Time) int64 {
	return time2epoch(
		t,
		big.NewInt(1),
		big.NewInt(978307200),
	)
}

// GoogleCalendar seems to count 32-day months from the day before the
// Unix epoch. @noppers worked out how to do this.
func GoogleCalendar(num int64) time.Time {

	n := int(num)

	totalDays := n / secondsPerDay
	seconds := n % secondsPerDay

	// A "Google month" has 32 days!
	months := totalDays / 32
	days := totalDays % 32

	// The "Google epoch" is apparently off by a day.
	t := time.Unix(-secondsPerDay, 0).UTC()

	// Add the days first...
	u := t.AddDate(0, 0, days)

	// ...then the months...
	v := u.AddDate(0, months, 0)

	// ...then the seconds.
	w := v.Add(time.Duration(seconds * 1e9))

	return w
}

// ToGoogleCalendar returns the GoogleCalendar time for the given time.Time.
func ToGoogleCalendar(t time.Time) int64 {
	y := t.Year() - 1970
	m := int(t.Month()) - 1
	r := ((((y*12+m)*32+t.Day())*24+t.Hour())*60+t.Minute())*60 + t.Second()
	return int64(r)
}

// ICQ time is the number of days since 1899-12-30, which is
// 2,209,161,600 seconds before the Unix epoch. Days can have a
// fractional part.
func ICQ(days float64) time.Time {

	t := time.Unix(-2209161600, 0).UTC()

	intdays := int(days)

	// Want the fractional part of the day in nanoseconds.
	fracday := int64((days - float64(intdays)) * nanosecondsPerDay)

	return t.AddDate(0, 0, intdays).Add(time.Duration(fracday))
}

// ToICQ returns the ICQ time for the given time.Time.
func ToICQ(t time.Time) float64 {
	t2 := time.Unix(-2209161600, 0)
	return float64(t.Sub(t2).Nanoseconds()) / float64(nanosecondsPerDay)
}

// Java time is the number of milliseconds since the Unix epoch.
func Java(num int64) time.Time {
	return epoch2time(
		big.NewInt(num),
		big.NewInt(1000),
		big.NewInt(0),
	)
}

// ToJava returns the Java time for the given time.Time.
func ToJava(t time.Time) int64 {
	return time2epoch(
		t,
		big.NewInt(1000),
		big.NewInt(0),
	)
}

// Mozilla time (e.g., formhistory.sqlite) is the number of
// microseconds since the Unix epoch.
func Mozilla(num int64) time.Time {
	return epoch2time(
		big.NewInt(num),
		big.NewInt(1e6),
		big.NewInt(0),
	)
}

// ToMozilla returns the Mozilla time for the given time.Time.
func ToMozilla(t time.Time) int64 {
	return time2epoch(
		t,
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

// ToOLE returns the OLE time for the given time.Time.
func ToOLE(t time.Time) string {
	icq := ToICQ(t)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, math.Float64bits(icq))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return fmt.Sprintf("%016x", buf.Bytes())
}

// Symbian time is the number of microseconds since the year 0, which
// is 62,167,219,200 seconds before the Unix epoch.
func Symbian(num int64) time.Time {
	return epoch2time(
		big.NewInt(num),
		big.NewInt(1e6),
		big.NewInt(-62167219200),
	)
}

// ToSymbian returns the Symbian time for the given time.Time.
func ToSymbian(t time.Time) int64 {
	return time2epoch(
		t,
		big.NewInt(1e6),
		big.NewInt(-62167219200),
	)
}

// Unix time is the number of seconds since 1970-01-01.
func Unix(num int64) time.Time {
	return time.Unix(num, 0).UTC()
}

// ToUnix returns the Unix time for the given time.Time.
func ToUnix(t time.Time) int64 {
	return t.Unix()
}

// UUIDv1 time (RFC 4122) is the number of hectonanoseconds (100 ns)
// since 1582-10-15, which is 12,219,292,800 seconds before the Unix
// epoch.
func UUIDv1(num int64) time.Time {
	return epoch2time(
		big.NewInt(num),
		big.NewInt(1e7),
		big.NewInt(-12219292800),
	)
}

// ToUUIDv1 returns the UUIDv1 time for the given time.Time.
func ToUUIDv1(t time.Time) int64 {
	return time2epoch(
		t,
		big.NewInt(1e7),
		big.NewInt(-12219292800),
	)
}

// WindowsDate time (e.g., .NET) is the number of hectonanoseconds
// (100 ns) since 0001-01-01, which is 62,135,596,800 seconds before
// the Unix epoch.
func WindowsDate(num int64) time.Time {
	return epoch2time(
		big.NewInt(num),
		big.NewInt(1e7),
		big.NewInt(-62135596800),
	)
}

// ToWindowsDate returns the WindowsDate time for the given time.Time.
func ToWindowsDate(t time.Time) int64 {
	return time2epoch(
		t,
		big.NewInt(1e7),
		big.NewInt(-62135596800),
	)
}

// WindowsFile time (e.g., NTFS) is the number of hectonanoseconds
// (100 ns) since 1601-01-01, which is 11,644,473,600 seconds before
// the Unix epoch.
func WindowsFile(num int64) time.Time {
	return epoch2time(
		big.NewInt(num),
		big.NewInt(1e7),
		big.NewInt(-11644473600),
	)
}

// ToWindowsFile returns the WindowsFile time for the given time.Time.
func ToWindowsFile(t time.Time) int64 {
	return time2epoch(
		t,
		big.NewInt(1e7),
		big.NewInt(-11644473600),
	)
}
