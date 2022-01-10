package energy

//NewDefaultEnergyWire ...
func NewDefaultEnergyWire() chan uint16 {
	return make(chan uint16, 3)
}
