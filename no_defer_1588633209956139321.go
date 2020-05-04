package main

import (
	"fmt"
	"os"
)

func aNewFunc() error {
	f, err := os.Open("/tmp/test.txt") //ISSUE
	if err != nil {
		return err
	}
	b := make([]byte, 5)
	n, err := f.Read(b)
	if err != nil {
		return err
	}
	fmt.Printf("%d bytes: %s\n", n, string(b))
	return nil
}
