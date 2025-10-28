package protos

type CafeGetInfoRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountServerId int64 `json:",omitempty,omitzero"`
}
