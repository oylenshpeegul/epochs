package epochs

import (
	"math/big"
	"time"
)

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
	return epoch(big.NewInt(num), big.NewInt(1e6), big.NewInt(-11644473600))
}

// Cocoa time is the number of seconds since 2001-01-01, which
// is 978,307,200 seconds after the Unix epoch.
func Cocoa(num int64) time.Time {
	return epoch(big.NewInt(num), big.NewInt(1), big.NewInt(978307200))
}

// Unix time is the number of seconds since 1970-01-01.
func Unix(num int64) time.Time {
	return time.Unix(num, 0).UTC()
}
