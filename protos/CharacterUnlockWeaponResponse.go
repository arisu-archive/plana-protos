package protos

type CharacterUnlockWeaponResponse struct {
	ResponsePacket
	WeaponDB WeaponDB `json:",omitempty,omitzero"`
}
