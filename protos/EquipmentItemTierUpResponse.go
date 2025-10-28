package protos

type EquipmentItemTierUpResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EquipmentDB EquipmentDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
}
