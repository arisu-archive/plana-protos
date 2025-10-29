package protos

type AccountBirthDayResponse struct {
	ResponsePacket
	AccountDB AccountDB `json:",omitempty,omitzero"`
}
