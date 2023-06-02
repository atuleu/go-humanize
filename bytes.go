// Package humanize provides utility type to format values in a
// readable manner. It focuse on concision over precision.
package humanize

import (
	"fmt"
	"math"
)

// A ByteSize represents a number of bytes, that formats itself using
// IEC binary prefixes with a radix of 2^10 (1024) like KiB, MiB
// GiB. It cannot go beyond a fe EbbiByte because we use int64.
type ByteSize int64

var iecPrefixes = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei"}

// String implement the fmt.Stringer interface
func (s ByteSize) String() string {
	value := float64(s)
	prefix := ""
	for _, prefix = range iecPrefixes {
		if math.Abs(value) < 1024.0 {
			break
		}
		value /= 1024.0
	}
	return fmt.Sprintf("%.1f %sB", value, prefix)
}
