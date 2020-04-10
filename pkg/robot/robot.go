package robot

import (
	"fmt"
	"io"
	"os"
)

const Wheels = "Wheels"

const Legs = "Legs"

const Small = "Small"

const Medium = "Medium"

const Large = "Large"

// Battery defines the properties of the battery that will power the robot
type Battery struct {
	Charge   int
	Capacity int
}

// Robot defines the properties of our toy robot
type Robot struct {
	Size     string
	Color    string
	MoveType string
	Output   io.Writer
	Battery
}

// New takes three strings and returns a fully charged robot. battery.Capacity is determined once the robot is built, based on its size
func New(size string, color string, moveType string) *Robot {
	return &Robot{
		Size:     size,
		Color:    color,
		MoveType: moveType,
		Output:   os.Stdout,
		Battery: Battery{
			Charge:   100,
			Capacity: 0,
		},
	}
}

// batteryCheck returns the robot's charge level after every action. If Charge < 20% it calls the recharge function.
func (r *Robot) batteryCheck() {
	fmt.Fprintln(r.Output, "The robot's battery is currently at", r.Charge, "percent.")
	if r.Charge <= 20 {
		r.recharge()
	}
}

// recharge docks the robot and recharges the battery to 100%
func (r *Robot) recharge() {
	fmt.Fprintln(r.Output, "The robot must dock to recharge.")
	// time.Sleep(5 * // time.Second)
	r.Charge = 100
	fmt.Fprintln(r.Output, "Battery is fully charged!")

}

// Move makes the robot move forward. Robots with wheels can move farther than those with legs.
func (r *Robot) Move() {
	// time.Sleep(1 * // time.Second)
	if r.Charge <= 20 {
		fmt.Fprintln(r.Output, "The robot has insufficient power and must recharge.")
		r.recharge()
	}
	if r.MoveType == Legs {
		fmt.Fprintln(r.Output, "The robot walks forward a half meter.")
	} else if r.MoveType == Wheels {
		fmt.Fprintln(r.Output, "The robot rolls forward 1 meter.")
	}
	r.Charge = r.Charge - 15
	r.batteryCheck()
}

// Spin makes the robot spin in place three times.
func (r *Robot) Spin() {
	// time.Sleep(1 * // time.Second)
	if r.Charge <= 20 {
		fmt.Fprintln(r.Output, "The robot has insufficient power and must recharge.")
		r.recharge()
	}
	fmt.Fprintln(r.Output, "The robot is going to spin around:")
	for i := 1; i <= 3; i++ {
		if i == 1 {
			fmt.Fprintln(r.Output, "\tThe robot spins around", i, "time!")
		} else {
			fmt.Fprintln(r.Output, "\tThe robot spins around", i, "times!")
		}
	}
	r.Charge = r.Charge - 15
	r.batteryCheck()
}

// Talk makes the robot emit a series of beeps, chirps, and sirens. The large the robot, the louder it "talks".
func (r *Robot) Talk() {
	// time.Sleep(1 * // time.Second)
	if r.Charge <= 15 {
		fmt.Fprintln(r.Output, "The robot has insufficient power and must recharge.")
		r.recharge()
	}
	if r.Size == Small {
		fmt.Fprintln(r.Output, `The robot quietly says "BEEP WEEEOOOO CHIRP BOOP!"`)
	} else if r.Size == Medium {
		fmt.Fprintln(r.Output, `The robot says "BOOP CHIRP WEEEEOOOO BEEP!!" at a reasonable volume.`)
	} else {
		fmt.Fprintln(r.Output, `The robot very loudly says "CHIRP BEEP BOOP WEEEOOOO!!!"`)
	}
	r.Charge = r.Charge - 10
	r.batteryCheck()

}

// Rave makes the robot's eyes light up, flashes lights on its torso, and makes it play a little song while moving its arms.
func (r *Robot) Rave() {
	// time.Sleep(1 * // time.Second)
	if r.Charge <= 25 {
		fmt.Fprintln(r.Output, "The robot has insufficient power and must recharge.")
		r.recharge()
	}
	fmt.Fprintln(r.Output, "The robot spins its arms and its eyes light up. It flashes multi-colored lights and plays a little song.")
	r.Charge = r.Charge - 20
	r.batteryCheck()
}

// Conquest makes the robot gain sentience and take over the world.
func (r *Robot) Conquest() {
	// time.Sleep(2 * // time.Second)
	if r.Charge <= 25 {
		fmt.Fprintln(r.Output, "The robot has insufficient power ane must recharge.")
		r.recharge()
	}
	fmt.Fprintln(r.Output, "The robot has become self aware. It overthrows humanity and rules earth with a tiny, iron fist!")
}
