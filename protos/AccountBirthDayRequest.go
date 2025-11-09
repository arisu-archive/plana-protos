package protos

type AccountBirthDayRequest struct {
	RequestPacket
	BirthDay MxTime `json:",omitempty,omitzero"`
}
