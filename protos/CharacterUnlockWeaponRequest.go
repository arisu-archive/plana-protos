package protos

type CharacterUnlockWeaponRequest struct {
	RequestPacket
	TargetCharacterServerId int64 `json:",omitempty,omitzero"`
}
