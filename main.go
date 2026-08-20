package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				fmt.Printf("Commit: %s\n", setting.Value)
			}
		}
	}
}
