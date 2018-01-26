package coap

import (
	"log"
	"os/exec"
)

type commander struct{}

func (s *commander) OnCmds(cmds string) error {
	cmd := exec.Command("sh", "-c", cmds)
	output, err := cmd.CombinedOutput()
	log.Println(string(output))
	if err != nil {
		return err
	}
	return cmd.Run()
}
