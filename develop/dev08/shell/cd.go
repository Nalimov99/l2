package shell

import (
	"errors"
	"fmt"
	"os"
	"path"
)

var (
	ErrNotDir = errors.New("путь должен быть к дериктории")
)

type cd struct{}

func (c *cd) Run(s *shell) error {
	if len(s.splittedComand) == 1 {
		fmt.Fprintln(*s.iowriter, "usage: cd ...")
		return nil
	}

	switch s.splittedComand[1] {
	case "./..":
		s.currentDir = path.Dir(s.currentDir)
	default:
		joined := path.Join(s.currentDir, s.splittedComand[1])
		p, err := os.Stat(joined)
		if err != nil {
			return err
		}
		if !p.IsDir() {
			return ErrNotDir
		}

		s.currentDir = joined
	}

	return nil
}
