package protos

type EquipmentChangePieceInfo struct {
	SourceItemId      int64 `json:",omitempty,omitzero"`
	ConsumeCount      int64 `json:",omitempty,omitzero"`
	DestinationItemId int64 `json:",omitempty,omitzero"`
}
