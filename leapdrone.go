package leapdrone

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/ardrone"
	"github.com/hybridgroup/gobot/platforms/leap"
)

func NewController(lAdapter *leap.LeapMotionAdaptor, aAdapter *ardrone.ArdroneAdaptor) *Controller {
	drone := ardrone.NewArdroneDriver(aAdapter, "Drone")
	l := leap.NewLeapMotionDriver(lAdapter, "leap")
	return &Controller{
		LAdapter: lAdapter,
		AAdapter: aAdapter,
		LDriver:  l,
		ADriver:  drone,
	}
}

// Controller defines all needed robot elements.
type Controller struct {
	LAdapter *leap.LeapMotionAdaptor
	LDriver  *leap.LeapMotionDriver
	AAdapter *ardrone.ArdroneAdaptor
	ADriver  *ardrone.ArdroneDriver
}

func (c *Controller) DroneRobot(worker func()) *gobot.Robot {
	return gobot.NewRobot("drone",
		[]gobot.Connection{c.AAdapter},
		[]gobot.Device{c.ADriver},
		worker,
	)
}

func (c *Controller) LeapRobot(worker func()) *gobot.Robot {
	// create a leap motion bot to track the event
	return gobot.NewRobot("leapBot",
		[]gobot.Connection{c.LAdapter},
		[]gobot.Device{c.LDriver},
		worker,
	)
}
