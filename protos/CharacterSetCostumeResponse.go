package protos

type CharacterSetCostumeResponse struct {
	ResponsePacket
	SetCostumeDB   CostumeDB
	UnsetCostumeDB CostumeDB
}
