package protos

type CraftBeginProcessResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CraftInfoDB CraftInfoDB `json:",omitempty,omitzero"`
}
