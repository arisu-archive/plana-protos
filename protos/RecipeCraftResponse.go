package protos

type RecipeCraftResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	EquipmentConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
	ItemConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
}
