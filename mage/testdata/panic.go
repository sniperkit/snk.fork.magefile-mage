/*
Sniperkit-Bot
- Status: analyzed
*/

// +build mage

package main

import (
	"errors"
)

// Function that panics.
func Panics() {
	panic("boom!")
}

// Error function that panics.
func PanicsErr() error {
	panic(errors.New("kaboom!"))
}
