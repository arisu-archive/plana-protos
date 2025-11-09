package protos

type ItemLockRequest struct {
	RequestPacket
	TargetServerId int64 `json:",omitempty,omitzero"`
	IsLocked       bool  `json:",omitempty,omitzero"`
}
