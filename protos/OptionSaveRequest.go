package protos

type OptionSaveRequest struct {
	RequestPacket
	OptionDB *OptionDB `json:",omitempty,omitzero"`
}
