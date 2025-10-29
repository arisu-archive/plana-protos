package protos

type CharacterListResponse struct {
	ResponsePacket
	CharacterDBs []CharacterDB `json:",omitempty,omitzero"`
	TSSCharacterDBs []CharacterDB `json:",omitempty,omitzero"`
	WeaponDBs []WeaponDB `json:",omitempty,omitzero"`
	CostumeDBs []CostumeDB `json:",omitempty,omitzero"`
}
