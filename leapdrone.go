package leapdrone

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/ardrone"
	"github.com/hybridgroup/gobot/platforms/leap"
)

// NewController returns a controller struct
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
// XXX: refactor by adding interface so we can have mock testing.
type Controller struct {
	LAdapter *leap.LeapMotionAdaptor
	LDriver  *leap.LeapMotionDriver
	AAdapter *ardrone.ArdroneAdaptor
	ADriver  *ardrone.ArdroneDriver
}

// DroneRobot returns an ARDrone robot.
func (c *Controller) DroneRobot(worker func()) *gobot.Robot {
	return gobot.NewRobot("drone",
		[]gobot.Connection{c.AAdapter},
		[]gobot.Device{c.ADriver},
		worker,
	)
}

// LeapRobot returns an Leap Motion robot.
func (c *Controller) LeapRobot(worker func()) *gobot.Robot {
	// create a leap motion bot to track the event
	return gobot.NewRobot("leapBot",
		[]gobot.Connection{c.LAdapter},
		[]gobot.Device{c.LDriver},
		worker,
	)
}
