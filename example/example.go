package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/ardrone"
	"github.com/hybridgroup/gobot/platforms/leap"
	"github.com/isarbit/leapdrone"
	lb "github.com/isarbit/leapdrone/leap"
)

func main() {
	gbot := gobot.NewGobot()

	// create a ardrone adapter
	aAdapter := ardrone.NewArdroneAdaptor("Drone")
	// create a leap adaptor to connect to the leap motion via web socket.
	lAdapter := leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
	c := leapdrone.NewController(lAdapter, aAdapter)

	// implement leap worker
	leapWorker := func() {
		// NOTE: leap motion driver only add message event to it.
		gobot.On(c.LDriver.Event("message"), func(data interface{}) {
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

	// implement ARDrone worker
	droneWorker := func() {
		drone.TakeOff()
		gobot.On(c.ADriver.Event("flying"), func(data interface{}) {
			gobot.After(3*time.Second, func() {
				drone.Land()
			})
		})
	}

	// add a leap robot to gobot
	gbot.AddRobot(c.LeapRobot(leapWorker))
	// add a ardrone robot to gobot
	gbot.AddEvent(c.DroneRobot(droneWorker))

	// start
	gbot.Start()
}
