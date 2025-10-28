package protos

type CharacterGearTierUpResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	GearDB GearDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
}
