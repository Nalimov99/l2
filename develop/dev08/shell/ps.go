package shell

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

type ps struct{}

func (p *ps) run(s *shell) error {
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
			// We only care if the name starts with a numeric
			if name[0] < '0' || name[0] > '9' {
				continue
			}

			// From this point forward, any errors we just ignore, because
			// it might simply be that the process doesn't exist anymore.
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

			fmt.Printf("%s %d\n", string(contents), int(pid))
		}
	}

	return nil
}
