package command

// OffCommand ...
type OffCommand struct {
	Device Device
}


// Execute ...
func (oc *OffCommand) Execute() {
	oc.Device.Off()
}