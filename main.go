package main

import (
	"runtime/debug"
)

// The import path of your module as defined in go.mod
const modulePath = "github.com/TimoG-WI/go-version-demo"

// Version returns the version of this library embedded in the binary.
func Version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	// 1. If consumed as a module dependency in another application:
	for _, dep := range info.Deps {
		if dep.Path == modulePath {
			return dep.Version
		}
	}

	// 2. If running directly from inside this module (e.g. tests or main binary):
	if info.Main.Path == modulePath && info.Main.Version != "" {
		return info.Main.Version
	}

	return "unknown"
}

func main(){
	fmt.Println("You are running version:", Version())
}
