package main

import "github.com/egor3f/loggo/cmd"

var version string

func main() {
	cmd.BuildVersion = version
	cmd.Initiate()
}
