package shell

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

type ps struct{}

func (p *ps) Run(s *shell) error {
	d, err := os.Open("/proc")
	if err != nil {
		return err
	}
	defer d.Close()

	for {
		names, err := d.Readdirnames(10)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		for _, name := range names {
			if name[0] < '0' || name[0] > '9' {
				continue
			}

			pid, err := strconv.ParseInt(name, 10, 0)
			if err != nil {
				continue
			}

			if err != nil {
				continue
			}

			contents, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", int(pid)))
			if err != nil {
				continue
			}

			fmt.Fprintf(*s.iowriter, "%s %d\n", string(contents), int(pid))
		}
	}

	return nil
}
