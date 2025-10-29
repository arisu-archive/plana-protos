package protos

type AccountCallNameResponse struct {
	ResponsePacket
	AccountDB AccountDB `json:",omitempty,omitzero"`
}
