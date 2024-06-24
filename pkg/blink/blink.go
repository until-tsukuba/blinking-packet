package blink

import ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"

// Sequence is an interface that represents a sequence of colors
type Sequence interface {
	Seq() []uint32
	Reversed() bool
}

var (
	ledCount int
	dev      *ws2811.WS2811
)

// SetUpWS2812 sets up the WS2812 LED strip
func SetUpWS2812(mLedCount int) {

	ledCount = mLedCount

	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = 255
	opt.Channels[0].LedCount = ledCount
	opt.Channels[0].GpioPin = 18

	dev, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		panic(err)
	}

	err = dev.Init()
	if err != nil {
		panic(err)
	}
}

// Blink sets the colors of the LEDs
func Blink(seq Sequence, diff int) {
	colors := seq.Seq()
	length := len(colors)
	reversed := seq.Reversed()

	for i := 0; i < ledCount; i++ {
		ledi := i
		if reversed {
			ledi = ledCount - i - 1
		}
		if i < diff || i >= length+diff {
			dev.Leds(0)[ledi] = 0
		} else {
			dev.Leds(0)[ledi] = colors[i-diff]
		}
	}

	err := dev.Render()
	if err != nil {
		panic(err)
	}
}

// Finish cleans up the WS2812 LED strip
// This function should be called at the end of the program
func Finish() {
	dev.Fini()
}
