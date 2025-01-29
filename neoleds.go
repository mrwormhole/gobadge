package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/shifter"
)

type NeoLedsUI struct{}

func (u NeoLedsUI) Start() {
	display.EnableBacklight(false)
	display.FillScreen(color.RGBA{0, 0, 0, 255})
	ledColors := make([]color.RGBA, 5)
	var i uint8
	for {
		ledColors[0] = getRainbowRGB(i)
		ledColors[1] = getRainbowRGB(i + 10)
		ledColors[2] = getRainbowRGB(i + 20)
		ledColors[3] = getRainbowRGB(i + 30)
		ledColors[4] = getRainbowRGB(i + 40)
		leds.WriteColors(ledColors)

		buttons.ReadInput()
		if buttons.Pins[shifter.BUTTON_START].Get() {
			break
		}
		i += 2

		time.Sleep(50 * time.Millisecond)
	}

	time.Sleep(50 * time.Millisecond)
	ledColors[0] = color.RGBA{0, 0, 0, 255}
	ledColors[1] = color.RGBA{0, 0, 0, 255}
	ledColors[2] = color.RGBA{0, 0, 0, 255}
	ledColors[3] = color.RGBA{0, 0, 0, 255}
	ledColors[4] = color.RGBA{0, 0, 0, 255}
	leds.WriteColors(ledColors)
	time.Sleep(50 * time.Millisecond)

	display.EnableBacklight(true)
}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 85 {
		return color.RGBA{i * 3, 255 - i*3, 0, 255}
	} else if i < 170 {
		i -= 85
		return color.RGBA{255 - i*3, 0, i * 3, 255}
	}
	i -= 170
	return color.RGBA{0, i * 3, 255 - i*3, 255}
}
