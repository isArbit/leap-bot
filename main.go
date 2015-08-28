package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/leap"
	lb "github.com/isarbit/leap-bot"
)

func main() {
	gbot := gobot.NewGobot()

	// create a leap adaptor to connect to the leap motion via web socket.
	leapMotionAdaptor := leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
	l := leap.NewLeapMotionDriver(leapMotionAdaptor, "leap")

	work := func() {
		// NOTE: leap motion driver only add message event to it.
		gobot.On(l.Event("message"), func(data interface{}) {
			lp := data.(leap.Frame)
			for _, v := range lp.Hands {
				// XXX: catch the first hand id and ignore other annoying hands
				hand := lb.NewHand(v, 200)
				if hand.IsForward() {
					fmt.Printf("moving forward\n")
				}
				if hand.IsBackward() {
					fmt.Printf("moving backward\n")
				}
				if hand.IsRight() {
					fmt.Printf("moving right\n")
				}
				if hand.IsLeft() {
					fmt.Printf("moving left\n")
				}
				if hand.IsUpward() {
					fmt.Printf("moving up\n")
				}
				if hand.IsDownward() {
					fmt.Printf("moving down\n")
				}
				continue
			}
		})
	}

	// create a leap motion bot to track the event
	robot := gobot.NewRobot("leapBot",
		[]gobot.Connection{leapMotionAdaptor},
		[]gobot.Device{l},
		work,
	)

	// add a leap robot to gobot
	gbot.AddRobot(robot)

	// start
	gbot.Start()
}
