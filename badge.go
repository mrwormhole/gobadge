package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"
)

type Badge struct {
	rainbow []color.RGBA
	quit    bool
}

func NewBadge() *Badge {
	rainbow := make([]color.RGBA, 256)
	for i := 0; i < 256; i++ {
		rainbow[i] = getRainbowRGB(uint8(i))
	}

	if Name == "" {
		Name = DefaultName
	}

	if Slogan1 == "" {
		Slogan1 = DefaultSlogan1
	}

	if Slogan2 == "" {
		Slogan2 = DefaultSlogan2
	}
	return &Badge{rainbow: rainbow}
}

func (b *Badge) Draw() {
	b.quit = false
	display.FillScreen(colors[GOPHERBLUE])

	for {
		b.logo()
		if b.quit {
			break
		}
		display.FillScreen(colors[WHITE])
		b.showSlogan(Slogan1, Slogan2)
		if b.quit {
			break
		}
		display.FillScreen(colors[GOPHERBLUE])
		b.scrollGoBanner()
		if b.quit {
			break
		}
		display.FillScreen(colors[WHITE])
		b.greetSelf(Name)
		if b.quit {
			break
		}
	}
}

func (b *Badge) logo() {
	const logoDisplayTime = 5 * time.Second
	display.FillRectangleWithBuffer(0, 0, WIDTH, HEIGHT, logoRGBA)
	time.Sleep(logoDisplayTime)
}

func (b *Badge) greetSelf(name string) {
	const r int16 = 8

	// black corners
	display.FillRectangle(0, 0, r, r, colors[BLACK])
	display.FillRectangle(0, HEIGHT-r, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, 0, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, HEIGHT-r, r, r, colors[BLACK])

	// round corners
	tinydraw.FilledCircle(&display, r, r, r, colors[GOPHERBLUE])
	tinydraw.FilledCircle(&display, WIDTH-r-1, r, r, colors[GOPHERBLUE])
	tinydraw.FilledCircle(&display, r, HEIGHT-r-1, r, colors[GOPHERBLUE])
	tinydraw.FilledCircle(&display, WIDTH-r-1, HEIGHT-r-1, r, colors[GOPHERBLUE])

	// top band
	display.FillRectangle(r, 0, WIDTH-2*r-1, r, colors[GOPHERBLUE])
	display.FillRectangle(0, r, WIDTH, 26, colors[GOPHERBLUE])

	// bottom band
	display.FillRectangle(r, HEIGHT-r-1, WIDTH-2*r-1, r+1, colors[GOPHERBLUE])
	display.FillRectangle(0, HEIGHT-2*r-1, WIDTH, r, colors[GOPHERBLUE])

	// gophers fonts
	tinyfont.WriteLine(&display, &gophers.Regular32pt, 10, 32, "BXYZWB", colors[WHITE])
	tinyfont.WriteLine(&display, &gophers.Regular32pt, 10, 110, "AGENIV", colors[GOPHERBLUE])

	w32, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, name)
	for i := 0; i < 230; i++ {
		tinyfont.WriteLineColors(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 72, name, b.rainbow[i:])

		buttons.ReadInput()
		if buttons.Pins[shifter.BUTTON_START].Get() {
			b.quit = true
			break
		}
	}
}

func (b *Badge) scrollGoBanner() {
	const top, mid, bottom = "KEEP CALM", "AND", " ON"

	w32top, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, top)
	w32middle, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, mid)
	w32bottomRight, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, bottom)
	w32bottomLeft, _ := tinyfont.LineWidth(&gophers.Regular32pt, "H")

	tinyfont.WriteLine(&display, &gophers.Regular32pt, (WIDTH-int16(w32middle))/2+8, 30, "P", colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32top))/2, 40, top, colors[WHITE])
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32middle))/2, 70, mid, colors[WHITE])
	tinyfont.WriteLine(&display, &gophers.Regular58pt, (WIDTH-int16(w32bottomRight+w32bottomLeft+45))/2, 115, "H", colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32bottomRight+w32bottomLeft+45))/2+45, 100, bottom, colors[WHITE])

	display.SetScrollArea(0, 0)
	for k := 0; k < 4; k++ {
		for i := int16(159); i >= 0; i-- {
			buttons.ReadInput()
			if buttons.Pins[shifter.BUTTON_START].Get() {
				b.quit = true
				break
			}
			display.SetScroll(i)
			time.Sleep(10 * time.Millisecond)
		}
	}
	display.SetScroll(0)
	display.StopScroll()
}

func (b *Badge) showSlogan(topline, bottomline string) {
	w32top, _ := tinyfont.LineWidth(&freesans.Regular9pt7b, topline)
	w32bottom, _ := tinyfont.LineWidth(&freesans.Regular9pt7b, bottomline)
	const cycles = 8

	for cycle := 0; cycle < cycles; cycle++ {
		for i := 0; i < 17; i++ {
			colorOffset := uint8((i * 16) % 256)

			tinyfont.WriteLine(&display, &freesans.Regular9pt7b, (WIDTH-int16(w32top))/2, 50, topline, getRainbowRGB(colorOffset))
			tinyfont.WriteLine(&display, &freesans.Regular9pt7b, (WIDTH-int16(w32bottom))/2, 80, bottomline, getRainbowRGB(colorOffset))

			buttons.ReadInput()
			if buttons.Pins[shifter.BUTTON_START].Get() {
				b.quit = true
				return
			}
			time.Sleep(20 * time.Millisecond)
		}
	}
}
