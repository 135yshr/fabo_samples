package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/tty.usbmodem1411")
	led := gpio.NewLedDriver(firmataAdaptor, "13")
	temp := aio.NewAnalogSensorDriver(firmataAdaptor, "0", 1*time.Second)

	work := func() {
		gobot.Every(1*time.Second, func() {
			x, err := temp.Read()
			if err != nil {
				fmt.Println(err)
				return
			}
			volt := convert(x, 0, 1024, 0, 5000)
			tempValue := convert(volt, 300, 1600, -30, 100)
			fmt.Println("temperature=", tempValue)
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led, temp},
		work,
	)

	robot.Start()
}

func convert(v int, inMin, inMax, outMin, outMax int) int {
	return (v-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}
