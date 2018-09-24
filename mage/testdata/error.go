/*
Sniperkit-Bot
- Status: analyzed
*/

// +build mage

package main

import (
	"errors"
)

// Returns a non-nil error.
func ReturnsNonNilError() error {
	return errors.New("bang!")
}
