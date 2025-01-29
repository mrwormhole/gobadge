package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/drivers/ws2812"
)

const (
	WIDTH  = 160
	HEIGHT = 128
)

const (
	BLACK = iota
	WHITE
	FUCHSIA
	GREEN
	RED
	GOPHERBLUE
)

var (
	display st7735.Device
	buttons shifter.Device
	leds    ws2812.Device
	buzzer  machine.Pin
	colors  = [...]color.RGBA{
		{0, 0, 0, 255},
		{255, 255, 255, 255},
		{206, 48, 98, 255},
		{0, 255, 0, 255},
		{255, 0, 0, 255},
		{1, 173, 216, 255},
	}
)

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})
	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.SCL_PIN, SDA: machine.SDA_PIN})

	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	buttons = shifter.NewButtons()
	buttons.Configure()

	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds = ws2812.New(neo)

	buzzer = machine.A0
	buzzer.Configure(machine.PinConfig{Mode: machine.PinOutput})

	speaker := machine.SPEAKER_ENABLE
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker.High()

	var sounds SoundsUI
	var neoleds NeoLedsUI

	for {
		switch menu() {
		case 0:
			NewBadge().Draw()
			break
		case 1:
			NewGame().Start()
			break
		case 2:
			neoleds.Start()
			break
		case 3:
			sounds.Start()
			break
		default:
			break
		}
		time.Sleep(1 * time.Second)
	}
}
