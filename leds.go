package main

import (
	"image/color"
	"machine"
	"time"
)

func Leds() {
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

		pressed, _ := buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}
		i += 2

		time.Sleep(50 * time.Millisecond)
	}

	ledColors[0] = color.RGBA{0, 0, 0, 255}
	ledColors[1] = color.RGBA{0, 0, 0, 255}
	ledColors[2] = color.RGBA{0, 0, 0, 255}
	ledColors[3] = color.RGBA{0, 0, 0, 255}
	ledColors[4] = color.RGBA{0, 0, 0, 255}
	leds.WriteColors(ledColors)
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
	tripled := i * 3
	switch {
	case i < 85:
		return color.RGBA{tripled, 255 - tripled, 0, 255}
	case i < 170:
		tripled -= 85
		return color.RGBA{255 - tripled, 0, i * 3, 255}
	default:
		tripled -= 170
		return color.RGBA{0, tripled, 255 - tripled, 255}
	}
}
