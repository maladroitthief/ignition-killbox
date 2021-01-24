package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

/* TODO:
- Button https://gobot.io/documentation/drivers/button/
- Relay https://gobot.io/documentation/drivers/relay/
- LCD https://gobot.io/documentation/drivers/ssd1306/
*/ //
func main() {
	serialPort := "/dev/ttyACM0"

	firmataAdaptor := firmata.NewAdaptor(serialPort)
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
