package main

import (
	"flag"
	"log"

	"github.com/mrwormhole/gobadge/cmd/logos"
)

const (
	gopherconUK22Logo = "./cmd/assets/gcuk-2022.jpg"
	gopherconUK23Logo = "./cmd/assets/gcuk-2023.jpg"
	tinygoLogo        = "./cmd/assets/tinygo.jpg"
	fathomLogo        = "./cmd/assets/fathom.jpg"
	civoLogo          = "./cmd/assets/civo.jpg"
)

func main() {
	conf := flag.String("conf", tinygoLogo, "Choose the conference logo you want to (e.g.: tinygo, gcuk22)")
	flag.Parse()

	c := confs()
	logo, ok := c[*conf]
	if !ok {
		log.Println("I do not have yet this conf in my catalog.")
		return
	}

	logos.GenerateLogoRGBAFile(logo)
}

func confs() map[string]string {
	return map[string]string{
		"gcuk22": gopherconUK22Logo,
		"gcuk23": gopherconUK23Logo,
		"tinygo": tinygoLogo,
		"fathom": fathomLogo,
		"civo":   civoLogo,
	}
}
