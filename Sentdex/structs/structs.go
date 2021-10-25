package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_var float64 = 1.9063

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	streeing_wheel int16
	top_speed_kmh  float64
}

func (c *car) top_speed_kpm() float64 {
	return c.top_speed_kmh / usixteenbitmax
}

func top_speed_kmh_func(c *car) float64 {
	return c.top_speed_kmh / usixteenbitmax
}

func (c *car) new_top_speed(new_speed float64) {
	c.top_speed_kmh = new_speed
}

func main() {
	myCar := car{gas_pedal: 233,
		brake_pedal:    0,
		streeing_wheel: 23,
		top_speed_kmh:  105,
	}
	fmt.Println(myCar.top_speed_kmh)
	fmt.Println(top_speed_kmh_func(&myCar))
	(&myCar).new_top_speed(305.0)
	fmt.Println(myCar.top_speed_kmh)
}
