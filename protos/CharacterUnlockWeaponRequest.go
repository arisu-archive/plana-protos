package protos

type CharacterUnlockWeaponRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetCharacterServerId int64 `json:",omitempty,omitzero"`
}
