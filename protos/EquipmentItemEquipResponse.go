package protos

type EquipmentItemEquipResponse struct {
	ResponsePacket
	CharacterDB CharacterDB `json:",omitempty,omitzero"`
	EquipmentDBs []EquipmentDB `json:",omitempty,omitzero"`
}
