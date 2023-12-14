package main

import (
	"github.com/Nathan-Leary/elsa/dev"
	"runtime"

	"github.com/Nathan-Leary/elsa/bundler"
	"github.com/Nathan-Leary/elsa/cmd"
	"github.com/Nathan-Leary/elsa/core"
)

func main() {
	runtime.LockOSThread()
	cmd.Execute(cmd.Elsa{
		Run:    core.Run,
		Bundle: bundler.BundleModule,
		Dev:    dev.RunDev,
	})
}
