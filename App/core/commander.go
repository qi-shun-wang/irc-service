package core

import (
	"strings"

	"github.com/go-cmd/cmd"
)

//Commander struct.
type Commander struct{}

//OnCmds method .
func (s *Commander) OnCmds(cmds string) (string, error) {

	c := cmd.NewCmd("sh", "-c", cmds)
	statusChan := c.Start()

	finalStatus := <-statusChan
	c.Stop()
	info := strings.Join(finalStatus.Stdout, " ")

	// log.Println(info)
	return info, nil
}
