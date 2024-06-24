package dump

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/until-tsukuba/blinking-packet/pkg/blink"
	"github.com/until-tsukuba/blinking-packet/pkg/parse"
	"net"
	"time"
)

type BlinkSequence struct {
	colors   []uint32
	reversed bool
}

func (b BlinkSequence) Seq() []uint32 {
	return b.colors
}

func (b BlinkSequence) Reversed() bool {
	return b.reversed
}

// DumpAndBlink captures packets from the given device and
// blinks the WS2812 LED strip according to the packet contents.
// The delay parameter specifies the delay between each blink in milliseconds,
// and recommended its value is 20.
func DumpAndBlink(device string, ledCount int, delay int) {
	blink.SetUpWS2812(ledCount)
	defer blink.Finish()

	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	interfaceInfo := interfaceInfoFrom(device)
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		parsedPacket := parse.Parse(packet, interfaceInfo)
		colors := colorSequence(parsedPacket)
		seq := BlinkSequence{colors, parsedPacket.Reversed}
		for i := 0; i < 60; i++ {
			blink.Blink(seq, i)
			time.Sleep(time.Millisecond * time.Duration(delay))
		}
	}
}

func interfaceInfoFrom(device string) net.Interface {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, iface := range ifaces {
		if iface.Name == device {
			return iface
		}
	}
	panic("interface not found")
}

func colorSequence(packetLayers parse.ParsedPacket) []uint32 {
	colors := make([]uint32, 0)
	for _, layer := range packetLayers.LayerList {
		r, g, b := parse.ProtocolToColor(layer.ProtocolName)
		for i := 0; i < 5; i++ {
			colors = append(colors, toLedColor(r, g, b))
		}
	}
	return colors
}

// toLedColor converts RGB values to a single uint32 value
// that can be used to set the color of a WS2812 LED.
// The RGB values should be in the range 0-7.
func toLedColor(r int, g int, b int) uint32 {
	return uint32(r)<<20 | uint32(g)<<12 | uint32(b)<<4
}

// fromLedColor converts a single uint32 value to RGB values.
// The RGB values will be in the range 0-7.
// This is the inverse of toLedColor.
func fromLedColor(led uint32) (r int, g int, b int) {
	r = int((led >> 20) & 0x7)
	g = int((led >> 12) & 0x7)
	b = int((led >> 4) & 0x7)
	return
}
