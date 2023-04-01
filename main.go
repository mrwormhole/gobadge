package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/drivers/ws2812"
)

var display st7735.Device
var buttons shifter.Device
var leds ws2812.Device
var bzrPin machine.Pin
var batterySensor machine.ADC

func main() {
	machine.InitADC()

	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})
	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.SCL_PIN, SDA: machine.SDA_PIN})

	batterySensor = machine.ADC{Pin: machine.A6} //14 592 -> 14 400
	batterySensor.Configure(machine.ADCConfig{})
	batterySensor.Get()

	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	buttons = shifter.NewButtons()
	buttons.Configure()

	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds = ws2812.New(neo)

	bzrPin = machine.A0
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	speaker := machine.SPEAKER_ENABLE
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker.High()

	for {
		switch menu() {
		case 0:
			NewBadge().Draw()
			break
		case 1:
			NewGame().Start()
			break
		case 2:
			Leds()
			break
		case 3:
			Music()
			break
		default:
			break
		}
		time.Sleep(1 * time.Second)
	}
}
