package command

// OnCommand ...
type OnCommand struct {
	Device Device
}


// Execute ...
func (oc *OnCommand) Execute() {
	oc.Device.On()
}