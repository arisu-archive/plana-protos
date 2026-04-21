package protos

type EquipmentItemEquipResponse struct {
	ResponsePacket
	CharacterDB              *CharacterDB `json:",omitempty,omitzero"`
	EquipmentDBs             []*EquipmentDB
	RemovedEquipmentServerId int64 `json:",omitempty,omitzero"`
}
