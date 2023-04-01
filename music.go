package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func Music() {
	white := color.RGBA{255, 255, 255, 255}
	display.FillScreen(white)

	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, 0, 50, "MUSIC", color.RGBA{0, 100, 250, 255})
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 20, 100, "Press any key", color.RGBA{200, 0, 0, 255})

	for {
		buttons.ReadInput()
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
			break
		}
		if buttons.Pins[shifter.BUTTON_START].Get() {
			tone(5274)
		}
		if buttons.Pins[shifter.BUTTON_A].Get() {
			tone(1046)
		}
		if buttons.Pins[shifter.BUTTON_B].Get() {
			tone(1975)
		}
		if buttons.Pins[shifter.BUTTON_LEFT].Get() {
			tone(329)
		}
		if buttons.Pins[shifter.BUTTON_UP].Get() {
			tone(369)
		}
		if buttons.Pins[shifter.BUTTON_DOWN].Get() {
			tone(523)
		}
		if buttons.Pins[shifter.BUTTON_RIGHT].Get() {
			tone(739)
		}
	}
}

func tone(tone int) {
	for i := 0; i < 10; i++ {
		bzrPin.High()
		time.Sleep(time.Duration(tone) * time.Microsecond)
		bzrPin.Low()
		time.Sleep(time.Duration(tone) * time.Microsecond)
	}
}
