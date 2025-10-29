package protos

type EquipmentItemTierUpResponse struct {
	ResponsePacket
	EquipmentDB EquipmentDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
}
