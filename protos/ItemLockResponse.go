package protos

type ItemLockResponse struct {
	ResponsePacket
	ItemDB ItemDB `json:",omitempty,omitzero"`
}
