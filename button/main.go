package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/tty.usbmodem1411")
	led := gpio.NewLedDriver(firmataAdaptor, "13")
	button := gpio.NewButtonDriver(firmataAdaptor, "12")

	work := func() {
		button.On("push", func(data interface{}) {
			led.Toggle()
		})
		button.On("release", func(data interface{}) {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led, button},
		work,
	)

	robot.Start()
}
