package times

import "time"

// durationを秒単位のIntにして返却
func DurationToSecondInt(duration time.Duration) int {
	return int(duration.Seconds())
}

// sleep by millisecond arg
func SleepMilliSecond(milli int) {
	time.Sleep(time.Millisecond * time.Duration(milli))
}
