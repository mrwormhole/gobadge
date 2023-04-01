package main

import (
	"time"

	"tinygo.org/x/drivers/shifter"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

func menu() int16 {
	bgColor := colors[FUCHSIA]
	options := []string{
		"Badge",
		"Snake",
		"NeoLeds",
		"Sounds",
		//"Settings", // disabled due to some goroutine funkyness
	}
	display.FillScreen(bgColor)

	selected := int16(0)
	numOpts := int16(len(options))
	for i := int16(0); i < numOpts; i++ {
		tinydraw.Circle(&display, 32, 37+10*i, 4, colors[WHITE])
		// applying shadow borders to fonts
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 39, 39+10*i, options[i], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 39, 40+10*i, options[i], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 39, 41+10*i, options[i], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40, 39+10*i, options[i], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40, 41+10*i, options[i], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41, 39+10*i, options[i], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41, 40+10*i, options[i], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41, 41+10*i, options[i], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40, 40+10*i, options[i], colors[GOPHERBLUE])
	}

	tinydraw.FilledCircle(&display, 32, 37, 2, colors[GOPHERBLUE])

	released := true
	for {
		pressed, _ := buttons.ReadInput()

		if released && buttons.Pins[shifter.BUTTON_UP].Get() && selected > 0 {
			selected--
			tinydraw.FilledCircle(&display, 32, 37+10*selected, 2, colors[GOPHERBLUE])
			tinydraw.FilledCircle(&display, 32, 37+10*(selected+1), 2, bgColor)
		}
		if released && buttons.Pins[shifter.BUTTON_DOWN].Get() && selected < (numOpts-1) {
			selected++
			tinydraw.FilledCircle(&display, 32, 37+10*selected, 2, colors[GOPHERBLUE])
			tinydraw.FilledCircle(&display, 32, 37+10*(selected-1), 2, bgColor)
		}
		if released && buttons.Pins[shifter.BUTTON_SELECT].Get() {
			break
		}
		if pressed == 0 {
			released = true
		} else {
			released = false
		}
		time.Sleep(200 * time.Millisecond)
	}
	return selected
}
