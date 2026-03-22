package protos

type RecipeCraftResponse struct {
	ResponsePacket
	ParcelResultDB           ParcelResultDB
	EquipmentConsumeResultDB ConsumeResultDB
	ItemConsumeResultDB      ConsumeResultDB
}
