package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

type Settings struct {
	SoundsEnabled, NeoLedsEnabled bool
}

func (s *Settings) Show() {
	bgColor := colors[FUCHSIA]
	options := []string{
		"Sounds",
		"NeoLeds",
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
		settings.DrawToggledText(options[i])
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
			switch selected {
			case 0:
				s.SoundsEnabled = !s.SoundsEnabled
				s.DrawToggledText("Sounds")
			case 1:
				s.NeoLedsEnabled = !s.NeoLedsEnabled
				s.DrawToggledText("NeoLeds")
			}
		}
		if released && buttons.Pins[shifter.BUTTON_START].Get() {
			break
		}
		if pressed == 0 {
			released = true
		} else {
			released = false
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (s *Settings) DrawToggledText(text string) {
	o, _ := tinyfont.LineWidth(&proggy.TinySZ8pt7b, text)
	offset := int16(o + 10)
	m := map[bool]string{true: "ON", false: "OFF"}
	c := map[string]color.RGBA{"ON": colors[GREEN], "OFF": colors[RED]}
	if text == "Sounds" {
		display.FillRectangle(int16(39+offset), 42-10, 20, 10, colors[FUCHSIA])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 39+offset, 40, m[s.SoundsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 39+offset, 41, m[s.SoundsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40+offset, 39, m[s.SoundsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40+offset, 41, m[s.SoundsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41+offset, 39, m[s.SoundsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41+offset, 40, m[s.SoundsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41+offset, 41, m[s.SoundsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, int16(40+offset), 40, m[s.SoundsEnabled], c[m[s.SoundsEnabled]])
	} else if text == "NeoLeds" {
		display.FillRectangle(int16(39+offset), 42, 20, 10, colors[FUCHSIA])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 39+offset, 40+10, m[s.NeoLedsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 39+offset, 41+10, m[s.NeoLedsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40+offset, 39+10, m[s.NeoLedsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40+offset, 41+10, m[s.NeoLedsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41+offset, 39+10, m[s.NeoLedsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41+offset, 40+10, m[s.NeoLedsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41+offset, 41+10, m[s.NeoLedsEnabled], colors[BLACK])
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, int16(40+offset), 40+10, m[s.NeoLedsEnabled], c[m[s.NeoLedsEnabled]])
	}
}
