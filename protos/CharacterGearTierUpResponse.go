package protos

type CharacterGearTierUpResponse struct {
	ResponsePacket
	GearDB          GearDB
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
}
