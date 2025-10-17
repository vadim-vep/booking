package main

import "testing"

func TestRun(t *testing.T) {
	_, err := run()
	if err != nil {
		panic(err)
	}
}
