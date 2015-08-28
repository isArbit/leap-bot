/*
	Package main now implements a leap motion moving.
*/
package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/leap"
)

// NewHand initiate a hand struct
func NewHand(h leap.Hand, th float64) *Hand {
	return &Hand{
		Hand:      h,
		ID:        h.ID,
		Threshold: th,
	}
}

// Hand is a extension of a leap motion hand by adding move direction
// detection.
type Hand struct {
	// embeded a leap motion hand
	leap.Hand

	ID        int
	Threshold float64
}

// IsUpward determines if hand is moving forward.
func (h *Hand) IsForward() bool {
	// negative Z-axis movement
	if h.PalmVelocity[2] < -(h.Threshold) {
		return true
	}
	return false
}

// IsUpward determines if hand is moving back.
func (h *Hand) IsBackward() bool {
	// positive Z-axis movement
	if h.PalmVelocity[2] > (h.Threshold) {
		return true
	}
	return false
}

// IsUpward determines if hand is moving up.
func (h *Hand) IsUpward() bool {
	// positive Y-axis movement
	if h.PalmVelocity[1] > (h.Threshold) {
		return true
	}
	return false
}

// IsDownward determines if hand is moving down.
func (h *Hand) IsDownward() bool {
	// negative Y-axis movement
	if h.PalmVelocity[1] < -(h.Threshold) {
		return true
	}
	return false
}

// IsRight determines if hand is moving right.
func (h *Hand) IsRight() bool {
	// positive X-axis movement
	if h.PalmVelocity[0] > (h.Threshold) {
		return true
	}
	return false
}

// IsLeft determines if hand is moving left.
func (h *Hand) IsLeft() bool {
	// negative X-axis movement
	if h.PalmVelocity[0] < -(h.Threshold) {
		return true
	}
	return false
}

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
				hand := NewHand(v, 200)
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
