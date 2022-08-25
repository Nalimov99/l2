package main

import (
	"dev08/shell"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestShellCommands(t *testing.T) {
	in := strings.NewReader("cd ./..\n")
	tmpfileout, err := ioutil.TempFile(".", "out")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(path.Join(".", tmpfileout.Name()))
	s, err := shell.StartShell(in, tmpfileout)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	cmds := s.Commands()
	_, ok := cmds["cd"]
	if !ok {
		t.Fatal("cd не найдена в командах")
	}
	_, ok = cmds["pwd"]
	if !ok {
		t.Fatal("pwd не найдена в командах")
	}
	_, ok = cmds["echo"]
	if !ok {
		t.Fatal("echo не найдена в командах")
	}
	_, ok = cmds["ps"]
	if !ok {
		t.Fatal("ps не найдена в командах")
	}
}

func TestShellCD(t *testing.T) {
	in := strings.NewReader("cd ./..\n")
	tmpfileout, err := ioutil.TempFile(".", "out")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(path.Join(".", tmpfileout.Name()))
	s, err := shell.StartShell(in, tmpfileout)
	if err != nil {
		t.Fatal(err)
	}

	p, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	want := path.Dir(p)
	got := s.CurrentDir()
	if want != got {
		t.Fatalf("want: %v, got: %v\n", want, got)
	}
}

func TestShellPWD(t *testing.T) {
	in := strings.NewReader("pwd..\n")
	tmpfileout, err := ioutil.TempFile(".", "out")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(path.Join(".", tmpfileout.Name()))
	s, err := shell.StartShell(in, tmpfileout)
	if err != nil {
		t.Fatal(err)
	}

	want, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	got := s.CurrentDir()
	if want != got {
		t.Fatalf("want: %v, got: %v\n", want, got)
	}
}
