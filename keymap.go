// Copyright 2020 Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/hypebeast/go-osc/osc"
	"github.com/micmonay/keybd_event"
)

type oscVal struct {
	addr string
	val  float32
}

type keyComb struct {
	key   int
	ctrl  bool
	alt   bool
	shift bool
}

var (
	keymap = map[oscVal]keyComb{
		oscVal{"/1/toggle1", 0}:        keyComb{keybd_event.VK_0, true, true, false}, // 방송중단
		oscVal{"/1/toggle1", 1}:        keyComb{keybd_event.VK_9, true, true, false}, // 방송시작
		oscVal{"/1/toggle2", 0}:        keyComb{keybd_event.VK_8, true, true, false}, // 녹화중단
		oscVal{"/1/toggle2", 1}:        keyComb{keybd_event.VK_7, true, true, false}, // 녹화시작
		oscVal{"/1/multipush2/1/1", 4}: keyComb{keybd_event.VK_1, true, true, false}, // 화면 1
		oscVal{"/1/multipush2/2/1", 4}: keyComb{keybd_event.VK_2, true, true, false}, // 화면 2
		oscVal{"/1/multipush2/3/1", 4}: keyComb{keybd_event.VK_1, true, true, false}, // 화면 3
		oscVal{"/1/multipush2/4/1", 4}: keyComb{keybd_event.VK_2, true, true, false}, // 화면 4
	}
)

func pressKeyComb(k keyComb) {
	kb.SetKeys(k.key)
	kb.HasCTRL(k.ctrl)
	kb.HasALT(k.alt)

	err := kb.Launching()
	chk(err)
}

func handlingOSCEvt(oscMsg *osc.Message) {
	oscV := oscVal{
		addr: oscMsg.Address,
		val:  oscMsg.Arguments[0].(float32),
	}

	if kc, ok := keymap[oscV]; ok {
		pressKeyComb(kc)
	}
}
