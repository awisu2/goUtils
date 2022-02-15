package times

import (
	"fmt"
	"time"
)

// durationを秒単位のIntにして返却
func DurationToSecondInt(duration time.Duration) int {
	return int(duration.Seconds())
}

// sleep by millisecond arg
func SleepMilliSecond(milli int) {
	fmt.Println("before")
	time.Sleep(time.Millisecond * time.Duration(milli))
	fmt.Println("after")
	fmt.Println("after1")
	fmt.Println("after2")
	time.Sleep(0)
	fmt.Println("after3")
	j := loop(0, 0)
	fmt.Println("after 4", j)
}

func loop(z int, j int) int {
	if z > 10 {
		return j
	}
	for i := 0; i < 10000000; i++ {
		j++
	}
	k := loop(z+1, j)
	return k
}
