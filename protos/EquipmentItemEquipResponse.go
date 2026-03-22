package protos

type EquipmentItemEquipResponse struct {
	ResponsePacket
	CharacterDB  CharacterDB
	EquipmentDBs []EquipmentDB
}
