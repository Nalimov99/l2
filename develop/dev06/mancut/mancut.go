package mancut

import (
	"errors"
	"os"
	"strings"
)

var (
	ErrFlagFInvalid = errors.New("поля нумеруются с 1")
)

type CutFlags struct {
	Field     int
	Delimiter string
	S         bool
}

type Mancut struct {
	Path  string
	lines []string
	CutFlags
}

func (m *Mancut) Result() (string, error) {
	if m.Field <= 0 {
		return "", ErrFlagFInvalid
	}

	if m.Delimiter == "" {
		m.Delimiter = "\t"
	}

	dat, err := os.ReadFile(m.Path)
	if err != nil {
		return "", err
	}

	m.lines = strings.Split(string(dat), "\n")
	return m.result(), nil
}

func (m *Mancut) result() string {
	res := make([]string, 0)

	for _, line := range m.lines {
		splittedLine := strings.Split(line, m.Delimiter)
		splittedLineLength := len(splittedLine)

		if splittedLineLength == 1 {
			if !m.S {
				res = append(res, line)
			}
			continue
		}

		if splittedLineLength < m.Field {
			res = append(res, "")
			continue
		}

		res = append(res, splittedLine[m.Field-1])
	}

	return strings.Join(res, "\n")
}
