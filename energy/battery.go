package energy

const (
	maxCapacity = 5000
	input       = 300
	output      = 300
)

type Accumulator struct {
	Charge uint64 //0-5000 kJ
}

//Charging give 300kW to accumulator charge
func (a *Accumulator) Charging() (isCharged bool) {
	if a.Charge != maxCapacity {
		a.Charge += input
		if a.Charge >= maxCapacity {
			a.Charge = maxCapacity
			isCharged = true
		}
	} else {
		isCharged = true
	}

	return isCharged
}

//Discharging take 300kW from accumulator charge
func (a *Accumulator) Discharging() (isPowerLow bool) {
	if a.Charge < output {
		a.Charge = 0
		isPowerLow = true
	} else {
		a.Charge -= output
	}
	return isPowerLow
}
