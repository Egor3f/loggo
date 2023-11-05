package main

import "github.com/aurc/loggo/cmd"

var version string

func main() {
	cmd.BuildVersion = version
	cmd.Initiate()
}
