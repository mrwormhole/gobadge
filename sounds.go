package main

import (
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

type SoundsUI struct{}

func (u SoundsUI) Start() {
	display.FillScreen(colors[WHITE])

	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, 0, 50, "MUSIC", colors[GOPHERBLUE])
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 20, 100, "Press any key", colors[FUCHSIA])

	for {
		buttons.ReadInput()
		if buttons.Pins[shifter.BUTTON_START].Get() {
			break
		}
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
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
		if buttons.Pins[shifter.BUTTON_RIGHT].Get() {
			tone(739)
		}
		if buttons.Pins[shifter.BUTTON_UP].Get() {
			tone(369)
		}
		if buttons.Pins[shifter.BUTTON_DOWN].Get() {
			tone(523)
		}
	}
}

// NOTE: this is only for gouroutine usage, do not use
func Sounds() {
	for {
		time.Sleep(3 * time.Millisecond)
		if !settings.SoundsEnabled {
			continue
		}

		buttons.ReadInput()
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
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
		time.Sleep(3 * time.Millisecond)
	}
}

func tone(tone int) {
	for i := 0; i < 10; i++ {
		buzzer.High()
		time.Sleep(time.Duration(tone) * time.Microsecond)
		buzzer.Low()
		time.Sleep(time.Duration(tone) * time.Microsecond)
	}
}
