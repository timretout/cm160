package cm160

import (
	"log"

	"github.com/kylelemons/gousb/usb"
)

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

var ctx *usb.Context

func init() {
	ctx = usb.NewContext()
}

const (
	CM160_VENDOR  = 0x0fde
	CM160_PRODUCT = 0xca05
)

func isCM160(desc *usb.Descriptor) bool {
	return desc.Vendor == usb.ID(CM160_VENDOR) &&
		desc.Product == usb.ID(CM160_PRODUCT)
}

// Open initializes and returns a device
func Open() *CM160 {
	devices, err := ctx.ListDevices(isCM160)
	if err != nil {
		log.Fatal(err)
	}
	if len(devices) == 0 {
		return nil
	}

	for i, v := range devices {
		if i != 0 {
			v.Close()
		}
	}

	d := devices[0]

	d.SetConfig(1)

	return &CM160{device: d}
}

func (c *CM160) Register(cb func(*Record)) {

}

func (c *CM160) Close() {
	c.device.Close()
}
