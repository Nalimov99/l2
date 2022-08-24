package shell

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type execCmd interface {
	run(*shell) error
}

type shell struct {
	currentDir     string
	reader         *bufio.Reader
	commands       map[string]execCmd
	currentCommand string
	splittedComand []string
}

func StartShell() error {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		return err
	}

	commands := make(map[string]execCmd)
	commands["cd"] = &cd{}
	commands["pwd"] = &pwd{}
	commands["echo"] = &echo{}
	commands["ps"] = &ps{}
	commands["kill"] = &kill{}

	s := shell{
		currentDir: currentDir,
		reader:     bufio.NewReader(os.Stdin),
		commands:   commands,
	}
	s.read()
	return nil
}

func (s *shell) read() {
	fmt.Print(">> ")
	text, err := s.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		defer s.read()
		return
	}

	s.currentCommand = text
	s.execute()
}

func (s *shell) execute() {
	s.currentCommand = strings.TrimSpace(s.currentCommand)
	s.currentCommand = strings.ToLower(s.currentCommand)
	whitespaces := regexp.MustCompile(`\s+`)
	s.currentCommand = whitespaces.ReplaceAllString(s.currentCommand, " ")

	if s.currentCommand == "\\quit" {
		return
	}

	s.splittedComand = strings.Split(s.currentCommand, " ")
	if runner, ok := s.commands[s.splittedComand[0]]; ok {
		if err := runner.run(s); err != nil {
			fmt.Println(err)
		}
	}

	s.read()
}
