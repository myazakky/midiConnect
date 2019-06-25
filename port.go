package midiConnect

import (
	"syscall"
)

type MidiInPort struct { 
	CallBackFunc uintptr
	Handle       uintptr
	winmm        *syscall.DLL
}

func NewMidiInPort(callBackFunc interface{}) (*MidiInPort, error) {
	winmm, err := syscall.LoadDLL("winmm")
	if err != nil {
		return nil, err
	}

	var handle uintptr

	m := new(MidiInPort)
	m.CallBackFunc = syscall.NewCallback(callBackFunc)
	m.Handle = handle
	m.winmm = winmm

	return m, nil
}
