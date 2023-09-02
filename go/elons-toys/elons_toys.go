package elon

import "fmt"

// Drive moves the car forward if it has enough battery to do so.
func (c *Car) Drive() {
	if c.battery > c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

// DisplayDistance returns a string indicating the total distance the car has driven.
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns a string indicating the remaining battery percent.
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish returns a bool indicating wether or not the car can finish a given distance.
func (c Car) CanFinish(trackDistance int) bool {
	return (c.battery/c.batteryDrain)*c.speed >= trackDistance
}
