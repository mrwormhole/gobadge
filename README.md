# GoBadge

My custom conference badge powered by TinyGo on top of Adafruit hardware ([EdgeBadge](https://www.adafruit.com/product/4400))

Important, in this fork, pressing "START" takes to main menu while "SELECT" directs to other apps. The main tinygo/gobadge repository uses the opposite.

# Compilation

- Run this command to compile and flash

If you are running Mac or Linux, or have [task](https://taskfile.dev/) installed you can run the following:

```
task flash
```

otherwise run tinygo directly

```
tinygo flash -target pybadge .
```

- To display a conference logo on your badge:
```
task flash CONF=gcuk22
```

- To customize the Gobadge with your own name and information, use the `NAME`, `SLOGAN1`, and `SLOGAN2` variables like this:

```
task flash NAME="@TinyGolang" SLOGAN1="Go compiler" SLOGAN2="small places"
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
		"customLogo"  : yourPathLogoHere,
	}
}
```

Add a new target to the Makefile:

```bash
	go run cmd/main.go -conf=customLogo
	tinygo flash -target pybadge .
```

You can run:

```bash
	task flash CONF=customLogo
```

It will use `cmd/logos/logo-template.txt` to generate the image into a `[]color.RGBA`.
Then it is stored in variable in `logo.go` file.

```go
package main

import "image/color"

var logoRGBA = []color.RGBA{ {255, 255, 255} }
```

After the image has been generated, the task command will flash it to the board.


👏 Congratulations! It is now a GoBadge.
