package protos

type AccountCallNameResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountDB AccountDB `json:",omitempty,omitzero"`
}
