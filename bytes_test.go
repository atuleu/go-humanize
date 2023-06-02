package humanize

import (
	"fmt"
	"math"
	"testing"
)

func TestBytesFormatting(t *testing.T) {
	testdata := []struct {
		Value    float64
		Expected string
	}{
		{Value: 0, Expected: "0.0 B"},
		{Value: 1023, Expected: "1023.0 B"},
		{Value: -10.4546 * math.Pow(2, 10), Expected: "-10.5 KiB"},
		{Value: 3.41 * math.Pow(2, 20), Expected: "3.4 MiB"},
		{Value: 9.0 * math.Pow(2, 30), Expected: "9.0 GiB"},
		{Value: 1.32 * math.Pow(2, 40), Expected: "1.3 TiB"},
		{Value: 7.96 * math.Pow(2, 50), Expected: "8.0 PiB"},
		{Value: 7.34 * math.Pow(2, 60), Expected: "7.3 EiB"},
	}

	for _, d := range testdata {
		formatted := fmt.Sprintf("%s", ByteSize(d.Value))
		if formatted != d.Expected {
			t.Errorf("Expected %d bytes to format to %s, got: %s",
				int64(d.Value), d.Expected, formatted)
		}
	}
}
