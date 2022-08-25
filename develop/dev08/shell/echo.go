package shell

import (
	"fmt"
	"strings"
)

type echo struct{}

func (e *echo) Run(s *shell) error {
	fmt.Fprint(*s.iowriter, strings.Join(s.splittedComand[1:], " "))

	return nil
}
