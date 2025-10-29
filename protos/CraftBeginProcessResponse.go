package protos

type CraftBeginProcessResponse struct {
	ResponsePacket
	CraftInfoDB CraftInfoDB `json:",omitempty,omitzero"`
}
