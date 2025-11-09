package protos

type EquipmentItemEquipRequest struct {
	RequestPacket
	CharacterServerId  int64   `json:",omitempty,omitzero"`
	EquipmentServerIds []int64 `json:",omitempty,omitzero"`
	EquipmentServerId  int64   `json:",omitempty,omitzero"`
	SlotIndex          int32   `json:",omitempty,omitzero"`
}
