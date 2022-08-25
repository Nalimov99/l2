package shell

import "fmt"

type pwd struct{}

func (p *pwd) Run(s *shell) error {
	fmt.Fprintln(*s.iowriter, s.currentDir)

	return nil
}
