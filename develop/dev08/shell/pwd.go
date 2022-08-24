package shell

import "fmt"

type pwd struct{}

func (p *pwd) run(s *shell) error {
	fmt.Println(s.currentDir)

	return nil
}
