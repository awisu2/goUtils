package times

import (
	"testing"
	"time"
)

func TestDurationToSecondInt(t *testing.T) {

	r := DurationToSecondInt(time.Duration(1 * time.Second))
	if r != 1 {
		t.Errorf("Invalid. not 1 %v", r)
	}

	r = DurationToSecondInt(time.Duration(1*time.Hour + 1*time.Minute + 1*time.Second))
	if r != 3661 {
		t.Errorf("Invalid. not 1 %v", r)
	}
}
