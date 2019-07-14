package midiConnect

import (
	"fmt"
	"strconv"
	"unsafe"
)

//dll methods
func (port *MidiInPort) Open(id int) error {
	proc, err := port.winmm.FindProc("midiInOpen")
	if err != nil {
		return err
	}
	proc.Call(
		uintptr(unsafe.Pointer(&port.Handle)),
		uintptr(id),
		port.CallBackFunc,
		0,
		0x00030000,
	)

	port.Start()

	return nil
}

func (port *MidiInPort) Start() error {
	proc, err := port.winmm.FindProc("midiInStart")
	if err != nil {
		return err
	}
	proc.Call(port.Handle)

	return nil
}

func (port *MidiInPort) Close() {
	proc, _ := port.winmm.FindProc("midiInClose")
	proc.Call(port.Handle)
}

func (port *MidiInPort) Stop() {
	proc, _ := port.winmm.FindProc("midiInStop")
	proc.Call(port.Handle)

	port.Close()
}

func GetStatus(wMsg, params int) (int, int, int) {
	if wMsg != 963 {
		return -1, -1, -1
	}
	hex := fmt.Sprintf("%x", params)

	switch len(hex) {
	case 5:
		hex = "0" + hex
	case 4:
		hex = "00" + hex
	}

	strength, _ := strconv.ParseInt(hex[0:2], 16, 32)
	tone, _ := strconv.ParseInt(hex[2:4], 16, 32)
	status, _ := strconv.ParseInt(hex[4:6], 16, 32)
	return int(strength), int(tone), int(status)

}
