// Copyright 2020 Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/hypebeast/go-osc/osc"
	"github.com/micmonay/keybd_event"
)

type oscVal struct {
	Addr string  `json: "addr"`
	Val  float32 `json: "val"`
}

type keyComb struct {
	Description string `jsono: description`
	Key         string `json: "key"`
	Ctrl        bool   `json: "ctrl"`
	Alt         bool   `json: "alt"`
	Shift       bool   `json: "shift"`
}

type keyShortcuts []struct {
	OSCVal  oscVal  `json: oscVal`
	KeyComb keyComb `json: keyComb`
}

var (
	shortcuts   *keyShortcuts
	shortcutMap = make(map[oscVal]keyComb)

	keymap = map[string]int{
		"0":   keybd_event.VK_0,
		"1":   keybd_event.VK_1,
		"2":   keybd_event.VK_2,
		"3":   keybd_event.VK_3,
		"4":   keybd_event.VK_4,
		"5":   keybd_event.VK_5,
		"6":   keybd_event.VK_6,
		"7":   keybd_event.VK_7,
		"8":   keybd_event.VK_8,
		"9":   keybd_event.VK_9,
		"F1":  keybd_event.VK_F1,
		"F2":  keybd_event.VK_F2,
		"F3":  keybd_event.VK_F3,
		"F4":  keybd_event.VK_F4,
		"F5":  keybd_event.VK_F5,
		"F6":  keybd_event.VK_F6,
		"F7":  keybd_event.VK_F7,
		"F8":  keybd_event.VK_F8,
		"F9":  keybd_event.VK_F9,
		"F10": keybd_event.VK_F10,
		"F11": keybd_event.VK_F11,
		"F12": keybd_event.VK_F12,
	}
)

func init() {
	var err error
	shortcuts, err = loadKeyShortcuts("shortcuts.json")
	chk(err)
	for _, s := range *shortcuts {
		shortcutMap[s.OSCVal] = s.KeyComb
	}
}

func pressKeyComb(k keyComb) {
	log.Println(k.Description)

	kb.SetKeys(keymap[k.Key])
	kb.HasCTRL(k.Ctrl)
	kb.HasALT(k.Alt)
	kb.HasSHIFT(k.Shift)

	// err := kb.Launching()
	// chk(err)

	kb.Press()
	time.Sleep(300 * time.Millisecond)
	kb.Release()
}

func handlingOSCEvt(oscMsg *osc.Message) {
	oscV := oscVal{
		Addr: oscMsg.Address,
		Val:  oscMsg.Arguments[0].(float32),
	}

	if kc, ok := shortcutMap[oscV]; ok {
		log.Println()
		pressKeyComb(kc)
	}
}

func loadKeyShortcuts(fn string) (*keyShortcuts, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ret keyShortcuts
	err = json.NewDecoder(f).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
