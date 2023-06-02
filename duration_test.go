package humanize

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	testdata := []struct {
		Value    time.Duration
		Expected string
	}{
		{Value: 0, Expected: "0s"},
		{Value: 1023, Expected: "1.023µs"},
		{Value: 29*time.Minute + 59*time.Second + 963*time.Millisecond, Expected: "29m59s"},
		{Value: 67*time.Minute + 18*time.Second, Expected: "1h7m"},
		{Value: 13*time.Hour + 58*time.Minute + 54*time.Second, Expected: "13h58m"},
		{Value: 58*time.Hour + 12*time.Minute, Expected: "2 days 10 hours"},
		{Value: 7*Day + 11*time.Hour, Expected: "1 week"},
		{Value: 9*Day + 13*time.Hour, Expected: "1 week 2 days"},
		{Value: 30*Day + 1*time.Hour, Expected: "1 month"},
		{Value: 5*Month + 3*Week + 2*Day, Expected: "5 months 3 weeks"},
		{Value: 12*Month + 1*Week + 2*Day, Expected: "1 year"},
		{Value: 26*Month + 3*Week + 2*Day, Expected: "2 years 2 months"},
	}

	for _, d := range testdata {
		formatted := fmt.Sprintf("%s", Duration(d.Value))
		if formatted != d.Expected {
			t.Errorf("Expected %s to format to %s, got: %s",
				d.Value, d.Expected, formatted)
		}
	}
}

func ExampleDuration() {
	fmt.Printf("%s\n", Duration(9*Day+45*time.Hour))
	// Output: 1 week 3 days
}
