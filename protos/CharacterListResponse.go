package protos

type CharacterListResponse struct {
	ResponsePacket
	CharacterDBs    []*CharacterDB
	TSSCharacterDBs []*CharacterDB
	WeaponDBs       []*WeaponDB
	CostumeDBs      []*CostumeDB
}
