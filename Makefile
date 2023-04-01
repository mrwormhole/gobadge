flash: flash-tinygo

flash-tinygo: prepare-tinygo perform-flash

flash-gcuk: prepare-gcuk perform-flash

prepare-tinygo:
	go run cmd/main.go -conf=tinygo

prepare-gcuk:
	go run cmd/main.go -conf=gcuk22

perform-flash:
	tinygo flash -size short -target pybadge -ldflags="-X main.YourName='$(NAME)' -X main.YourTitle1='$(TITLE1)' -X main.YourTitle2='$(TITLE2)'" .
