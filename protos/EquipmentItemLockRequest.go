package protos

type EquipmentItemLockRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetServerId int64 `json:",omitempty,omitzero"`
	IsLocked bool `json:",omitempty,omitzero"`
}
