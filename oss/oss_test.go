package oss

import (
	"fmt"
	"testing"
)

func TestReadStdOut(t *testing.T) {
	got, _ := ReadStdOut(func() {
		fmt.Println("hello world")
	})
	if string(got) != "hello world\n" {
		t.Errorf("not get hello world")
	}
}
