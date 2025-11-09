package protos

type RecipeCraftResponse struct {
	ResponsePacket
	ParcelResultDB           ParcelResultDB  `json:",omitempty,omitzero"`
	EquipmentConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
	ItemConsumeResultDB      ConsumeResultDB `json:",omitempty,omitzero"`
}
