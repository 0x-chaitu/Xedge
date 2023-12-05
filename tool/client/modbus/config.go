package modbus

type CommandInfo struct {
	DataBank        string
	StartingAddress uint16

	//total address to read from modbus
	TotalAddress uint16
}
