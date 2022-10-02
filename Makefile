default:
	go build -o ./bin/zigbee2webthings ./cmd/zigbee2webthings/main.go


rpi_build:
	env GOOS=linux GOARCH=arm GOARM=6 go build -o ./bin/zigbee2webthings ./cmd/zigbee2webthings/main.go

build_all:
	go build ./...