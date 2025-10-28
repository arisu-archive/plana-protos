package protos

type CraftInfoListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
