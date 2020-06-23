package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/hypebeast/go-osc/osc"
	"github.com/micmonay/keybd_event"
)

const (
	addr = "192.168.219.159:8765"
)

var (
	kb keybd_event.KeyBonding
)

func main() {
	var err error
	kb, err = keybd_event.NewKeyBonding()
	chk(err)
	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	server := &osc.Server{}
	conn, err := net.ListenPacket("udp", addr)
	chk(err)
	defer conn.Close()

	fmt.Println("Start listening on", addr)
	go touchOSC(server, conn)

	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, os.Interrupt)
	<-quitCh
	os.Exit(0)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
