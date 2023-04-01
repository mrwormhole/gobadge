# GoBadge

My custom conference badge powered by TinyGo on top of Adafruit hardware ([EdgeBadge](https://www.adafruit.com/product/4400))

# Compilation

- Run this command to compile and flash

If you are running Mac or Linux, or have make installed you can run the following:

```
make flash
```

otherwise run tinygo directly

```
tinygo flash -target pybadge .
```

- To display a conference logo on your badge:
```
make flash-gcuk
```

- To customize the Gobadge with your own name and information, use the `NAME`, `TITLE1`, and `TITLE2` variables like this:

```
make flash-gcus NAME="@TinyGolang" TITLE1="Go compiler" TITLE2="small places"
```

# Custom Logo

- Create an image with a 160x128 pixels size, copy it into `cmd/assets` folder.  
For the moment only jpeg images are supported.  
- In `cmd/main.go` add the path to your file here

```go
const (
    gopherconUK22Logo = "./cmd/assets/gopherconuk-2022.jpg"
    yourPathLogoHere = "./your/path/to/the/logo"
)
```

- Add the corresponding flag to the conf map:

```go
func confs() map[string]string {
	return map[string]string{
		"gcuk22"    : gopherconUK22Logo,
		"flagLogo"  : yourPathLogoHere,
	}
}
```

Add a new target to the Makefile:

```bash
flash-yourconf:
	go run cmd/main.go -conf=flagLogo
	tinygo flash -target pybadge .
```

You can run:

```bash
make flash-yourconf
```

It will use `cmd/logos/logo-template.txt` to generate the image into a `[]color.RGBA`.
Then it is stored in variable in `logo.go` file.

```go
package main

import "image/color"

var logoRGBA = []color.RGBA{ {255, 255, 255} }
```

After the image has been generated, the make command will flash it to the board.


👏 Congratulations! It is now a GoBadge.
