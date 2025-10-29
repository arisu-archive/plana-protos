package protos

type CafeGetInfoRequest struct {
	RequestPacket
	AccountServerId int64 `json:",omitempty,omitzero"`
}
