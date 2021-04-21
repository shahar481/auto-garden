package web_server

import (
	"auto-garden/testutility"
	"fmt"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"os"
	"testing"
)

func TestCreateListener(t *testing.T) {
	port, err := freeport.GetFreePort()
	if err != nil {
		t.Fatal(err)
	}
	args := testutility.CreateArgs(
		"",
		0,
		"",
		"",
		"",
		"127.0.0.1",
		port,
		"",
		"",
		false,
		"")
	l := createListener(args)
	defer l.Close()
	_, err = net.Dial("tcp", fmt.Sprintf("%s:%d", *(args.WebListenIp), *(args.WebListenPort)))
	if err != nil {
		t.Error(err)
	}
}

func getRandomFilePath(t *testing.T) string {
	f, err := ioutil.TempFile(".", "")
	if err != nil {
		t.Fatal(err)
	}
	path := f.Name()
	f.Close()
	err = os.Remove(path)
	if err != nil {
		t.Fatal(err)
	}
	return path
}

func TestCreateTlsListener(t *testing.T) {
	keyPath := getRandomFilePath(t)
	crtPath := getRandomFilePath(t)
	port, err := freeport.GetFreePort()
	if err != nil {
		t.Fatal(err)
	}
	args := testutility.CreateArgs(
		"",
		0,
		"",
		"",
		"",
		"127.0.0.1",
		port,
		keyPath,
		crtPath,
		true,
		"")
	assert.Panics(t, func(){createListener(args)})
}
