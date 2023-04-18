package main

import (
	"flag"
	"log"

	"github.com/tinygo-org/gobadge/cmd/logos"
)

const (
<<<<<<< Updated upstream
	gopherconEU22Logo = "./cmd/assets/gopherconeu-2022.jpg"
	gopherconUK22Logo = "./cmd/assets/gopherconuk-2022.jpg"
	gopherconUS22Logo = "./cmd/assets/gopherconus-2022.jpg"
	fosdem23Logo      = "./cmd/assets/fosdem-2023.jpg"
=======
	gopherconUK22Logo = "./cmd/assets/gcuk-2022.jpg"
	gopherconUK23Logo = "./cmd/assets/gcuk-2023.jpg"
>>>>>>> Stashed changes
	tinygoLogo        = "./cmd/assets/tinygo.jpg"
	fathomLogo        = "./cmd/assets/fathom.jpg"
)

func main() {
	conf := flag.String("conf", tinygoLogo, "Choose the conference logo you want to (e.g.: tinygo, gceu22, gcuk22, gcus22)")
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
<<<<<<< Updated upstream
		"gceu22":   gopherconEU22Logo,
		"gcuk22":   gopherconUK22Logo,
		"gcus22":   gopherconUS22Logo,
		"fosdem23": fosdem23Logo,
		"tinygo":   tinygoLogo,
=======
		"gcuk22": gopherconUK22Logo,
		"gcuk23": gopherconUK23Logo,
		"tinygo": tinygoLogo,
		"fathom": fathomLogo,
>>>>>>> Stashed changes
	}
}
