/*
Sniperkit-Bot
- Status: analyzed
*/

// +build mage

package main

import (
	"fmt"
	"log"

	"github.com/sniperkit/snk.fork.magefile-mage/mg"
)

// This should work as a default - even if it's in a different file
var Default = ReturnsNilError

// this should not be a target because it returns a string
func ReturnsString() string {
	fmt.Println("more stuff")
	return ""
}

func TestVerbose() {
	log.Println("hi!")
}

func ReturnsVoid() {
	mg.Deps(f)
}

func f() {}
