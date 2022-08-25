package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

type execCmd interface {
	Run(*shell) error
}

type shell struct {
	currentDir     string
	reader         *bufio.Reader
	commands       map[string]execCmd
	currentCommand string
	splittedComand []string
	ioreader       *io.Reader
	iowriter       *io.Writer
}

func StartShell(ioreader io.Reader, iowriter io.Writer) (*shell, error) {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}

	commands := make(map[string]execCmd)
	commands["cd"] = &cd{}
	commands["pwd"] = &pwd{}
	commands["echo"] = &echo{}
	commands["ps"] = &ps{}
	commands["kill"] = &kill{}

	s := shell{
		currentDir: currentDir,
		reader:     bufio.NewReader(ioreader),
		commands:   commands,
		ioreader:   &ioreader,
		iowriter:   &iowriter,
	}
	go s.stopShell()

	s.read()
	return &s, nil
}

func (s *shell) stopShell() {
	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt, syscall.SIGTERM)
	for {
		<-q
		fmt.Fprintln(*s.iowriter, "\nЕсли хотите завершить shell пропишите: \\quit")
	}
}

func (s *shell) SetCurrentDir(dir string) {
	s.currentDir = dir
}

func (s *shell) Commands() map[string]execCmd {
	return s.commands
}

func (s *shell) CurrentDir() string {
	return s.currentDir
}

func (s *shell) read() {
	fmt.Fprint(*s.iowriter, ">> ")
	text, err := s.reader.ReadString('\n')
	if err != nil {
		fmt.Fprint(*s.iowriter, err)
		return
	}

	s.currentCommand = text
	s.Execute()
}

func (s *shell) Execute() {
	s.currentCommand = strings.TrimSpace(s.currentCommand)
	s.currentCommand = strings.ToLower(s.currentCommand)
	whitespaces := regexp.MustCompile(`\s+`)
	s.currentCommand = whitespaces.ReplaceAllString(s.currentCommand, " ")

	if s.currentCommand == "\\quit" {
		return
	}

	s.splittedComand = strings.Split(s.currentCommand, " ")
	if runner, ok := s.commands[s.splittedComand[0]]; ok {
		if err := runner.Run(s); err != nil {
			fmt.Fprintln(*s.iowriter, err)
		}
	}

	s.read()
}
