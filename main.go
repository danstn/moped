package main

import "github.com/danstn/moped/cmd"

func main() {
	mopedCLI := cmd.NewMopedCLI()
	mopedCLI.Execute()
}
