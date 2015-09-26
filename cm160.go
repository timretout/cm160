package cm160

import "github.com/kylelemons/gousb/usb"

type CM160 struct {
	device *usb.Device
}

type Record struct {
	Volt   int
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Cost   float32
	Amps   float32
	IsLive bool
	Watt   float32
}

func Register(cb func(*Record)) {
}
