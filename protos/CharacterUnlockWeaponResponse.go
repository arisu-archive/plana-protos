package protos

type CharacterUnlockWeaponResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	WeaponDB WeaponDB `json:",omitempty,omitzero"`
}
