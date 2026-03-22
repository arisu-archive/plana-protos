package protos

type CharacterGearUnlockResponse struct {
	ResponsePacket
	GearDB      GearDB
	CharacterDB CharacterDB
}
