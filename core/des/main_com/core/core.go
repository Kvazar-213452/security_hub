package core

import (
	"head/main_com/module"
	"os/exec"
)

var Cmd *exec.Cmd
var CleanupDone = make(chan struct{})

func Cleanup() {
	if Cmd != nil && Cmd.Process != nil {
		Cmd.Process.Kill()
	}
	module.KillAllModules()
	close(CleanupDone)
}
