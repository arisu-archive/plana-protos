package protos

type ItemLockResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ItemDB ItemDB `json:",omitempty,omitzero"`
}
