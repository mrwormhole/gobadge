package main

import (
	"flag"
	"fmt"

	"github.com/mrwormhole/gobadge/cmd/logos"
)

const (
	gopherconUK22Logo = "./cmd/assets/gopherconuk-2022.jpg"
	tinygoLogo        = "./cmd/assets/tinygo.jpg"
)

func main() {
	conf := flag.String("conf", tinygoLogo, "Choose the conference logo you want to (e.g.: tinygo, gceu22, gcuk22, gcus22)")
	flag.Parse()

	c := confs()
	logo, ok := c[*conf]
	if !ok {
		fmt.Println("I do not have yet this conf in my catalog.")
		return
	}

	logos.GenerateLogoRGBAFile(logo)
}

func confs() map[string]string {
	return map[string]string{
		"gcuk22": gopherconUK22Logo,
		"tinygo": tinygoLogo,
	}
}
