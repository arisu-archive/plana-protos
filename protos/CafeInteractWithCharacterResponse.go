package protos

type CafeInteractWithCharacterResponse struct {
	ResponsePacket
	CafeDB         CafeDB
	CharacterDB    CharacterDB
	ParcelResultDB ParcelResultDB
}
