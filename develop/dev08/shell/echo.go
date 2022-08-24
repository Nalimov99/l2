package shell

import (
	"fmt"
	"strings"
)

type echo struct{}

func (e *echo) run(s *shell) error {
	fmt.Println(strings.Join(s.splittedComand[1:], " "))

	return nil
}
