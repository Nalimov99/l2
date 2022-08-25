package shell

import (
	"fmt"
	"os"
	"strconv"
)

type kill struct{}

func (k *kill) Run(s *shell) error {
	pid, err := strconv.Atoi(s.splittedComand[1])
	if err != nil {
		return err
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprint(*s.iowriter, err)
	}
	proc.Kill()
	return nil
}
