// Copyright 2020 Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/hypebeast/go-osc/osc"
)

func touchOSC(server *osc.Server, conn net.PacketConn) {

	for {
		packet, err := server.ReceivePacket(conn)
		if err != nil {
			fmt.Println("Server error: " + err.Error())
			os.Exit(1)
		}

		if packet != nil {
			switch packet.(type) {
			default:
				fmt.Println("Unknow packet type!")

			case *osc.Message:
				fmt.Printf("-- OSC Message: ")
				msg := packet.(*osc.Message)
				log.Println(msg)
				handlingOSCEvt(msg)
				// log.Println(msg.Address)
				// for _, arg := range msg.Arguments {
				// 	log.Println("   ", arg.(float32))
				// }
				// log.Println(pkt.Address, pkt.
				// osc.PrintMessage(packet.(*osc.Message))
			}
		}
	}
}
