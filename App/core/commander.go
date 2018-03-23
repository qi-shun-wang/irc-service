package core

import (
	"log"
	"strings"
	"time"

	"github.com/go-cmd/cmd"
)

//Commander struct.
type Commander struct{}

//OnDebugCmds method .
func (s *Commander) OnDebugCmds(cmds string) (string, error) {

	c := cmd.NewCmd("sh", "-c", cmds)
	statusChan := c.Start()

	var gotStatus cmd.Status
	var info string
	timeout := time.After(5 * time.Second)

	select {
	case gotStatus = <-statusChan:

		// a := []string{gotStatus.Error.Error(), gotStatus.St}
		log.Println(gotStatus)
		info = strings.Join(gotStatus.Stderr, ";")
		info = info + strings.Join(gotStatus.Stdout, ";")
		info = info + "=Cmd string=>" + gotStatus.Cmd + ";"
		info = info + "=exit code=>" + string(gotStatus.Exit)
	case <-timeout:
		c.Stop()
		info = "timeout waiting for statusChan"
	}

	return info, nil
}

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
