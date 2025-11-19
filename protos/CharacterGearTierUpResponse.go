package protos

type CharacterGearTierUpResponse struct {
	ResponsePacket
	GearDB          GearDB          `json:",omitempty,omitzero"`
	ParcelResultDB  ParcelResultDB  `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
}
