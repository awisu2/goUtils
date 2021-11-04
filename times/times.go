package times

import "time"

// durationを秒単位のIntにして返却
func DurationToSecondInt(duration time.Duration) int {
	return int(duration.Seconds())
}
