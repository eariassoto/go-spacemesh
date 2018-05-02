//go-spacemesh is a golang implementation of the Spacemesh node.
//See - https://spacemesh.io
package main

import (
	"github.com/spacemeshos/go-spacemesh/app"
	// "github.com/spacemeshos/go-spacemesh/cmd"
)

// vars set by make from outside
var (
	commit  = ""
	branch  = ""
	version = "0.0.1"
)

func main() { // run the app
	app.Main(commit, branch, version)
	// cmd.Execute()
}
