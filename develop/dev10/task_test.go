package main

import (
	"dev10/client"
	"dev10/server"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTelnet(t *testing.T) {
	in := strings.NewReader("123\n\n\n")
	in.Seek(0, 0)
	tmp, err := ioutil.TempFile(".", "out")
	if err != nil {
		t.Fatal(err)
	}
	defer tmp.Close()

	t.Cleanup(func() {
		os.Remove(tmp.Name())
	})
	ts := server.TelnetServer{}
	go ts.RunServer()
	client.NewTelnet("127.0.0.1:8020", time.Second*10, in, tmp)

	got, err := ioutil.ReadFile(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}

	want := ">> message recieved: 123"

	if string(got) != want {
		t.Fatalf("want: %s; got: %v", want, string(got))
	}
}
