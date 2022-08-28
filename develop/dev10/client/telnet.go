package client

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	ErrCouldNotConnect = errors.New("не возможно подключиться к серверу")
)

type telnet struct {
	address    string
	connection net.Conn
	ctx        context.Context
	quit       chan struct{}
	timeout    time.Duration
	in         io.Reader
	out        io.Writer
}

func NewTelnet(address string, timeout time.Duration, in io.Reader, out io.Writer) error {
	quit := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	t := telnet{
		address: address,
		ctx:     ctx,
		quit:    quit,
		timeout: timeout,
		in:      in,
		out:     out,
	}

	if err := t.runDial(); err != nil {
		cancel()
		return err
	}
	defer t.connection.Close()

	t.close()
	cancel()

	return nil
}

func (t *telnet) runDial() error {
	end := time.After(t.timeout)
	for {
		select {
		case <-end:
			return ErrCouldNotConnect
		default:
			con, err := net.Dial("tcp", t.address)
			if err != nil {
				fmt.Println(err)
				time.Sleep(time.Second)
				continue
			}
			t.connection = con
			go t.write()
			go t.read()
			return nil
		}
	}
}

func (t *telnet) write() {
	reader := bufio.NewReader(t.in)
	fmt.Print(">> ")

	for {
		select {
		case <-t.ctx.Done():
			return
		default:
			msg, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					t.quit <- struct{}{}
					return
				}
				continue
			}

			_, err = t.connection.Write([]byte(msg))
			if err != nil {
				fmt.Println(err)
				t.quit <- struct{}{}
			}
		}
	}
}

func (t *telnet) read() {
	reader := bufio.NewReader(t.connection)
	for {
		select {
		case <-t.ctx.Done():
			return
		default:
			msg, err := reader.ReadString('\n')
			if err != nil {
				if err = t.runDial(); err != nil {
					t.quit <- struct{}{}
				}

				return
			}

			msg = strings.TrimSuffix(msg, "\n")

			t.out.Write([]byte(">> " + msg))
			fmt.Print("\n>> ")
		}
	}
}

func (t *telnet) close() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-stop:
			fmt.Println("\nчтобы завершить программу используйте Ctrl+D")
		case <-t.quit:
			fmt.Println("\nВыходим")
			return
		}
	}
}
