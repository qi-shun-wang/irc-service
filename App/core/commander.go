package core

import (
	"log"
	"os/exec"
)

//Commander struct.
type Commander struct{}

//OnCmds method .
func (s *Commander) OnCmds(cmds string) error {
	cmd := exec.Command("sh", "-c", cmds)
	output, err := cmd.CombinedOutput()
	log.Println(string(output))
	if err != nil {
		return err
	}
	return cmd.Run()
}
