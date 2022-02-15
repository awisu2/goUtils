package output

import (
	"fmt"
	"os"
)

// output the set value to the standard output as it is
func Output(stdout string, stderr string) {
	if stdout != "" {
		fmt.Println(stdout)
	}
	if stderr != "" {
		fmt.Fprintln(os.Stderr, stderr)
	}
}
