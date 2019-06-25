midiConnect
====

midiConnect is a midi receiver library for Golang. 

## Platform
windows

## Usage
```go
package main

import (
	"github.com/nimuC/midiConnect"
	"fmt"
	"log"
)

func main() {

	f := func(hMidiIn uintptr, wMsg int, dwInstance uintptr, dwParam1 int, dwParam2 int) uintptr {

		strength, tone, status := midiConnect.GetStatus(wMsg, dwParam1)
		if tone != -1 {
			fmt.Println(strength, tone, status)
		}

		return 0
	}

	m, err := midiConnect.NewMidiInPort(f)
	if err != nil {
		log.Fatal(err)
	}

	m.Open(0)
	defer m.Close()

	for {
	}

}
```

## Install
`go get github.com/nimuC/midiConnect/`

## Licence
MIT
