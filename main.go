package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	fmt.Println(buildInfo.Settings)
}
