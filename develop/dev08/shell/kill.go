package shell

import (
	"log"
	"os"
	"strconv"
)

type kill struct{}

func (k *kill) run(s *shell) error {
	pid, err := strconv.Atoi(s.splittedComand[1])
	if err != nil {
		return err
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		log.Println(err)
	}
	proc.Kill()
	return nil
}
