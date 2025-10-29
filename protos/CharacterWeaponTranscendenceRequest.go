package protos

type CharacterWeaponTranscendenceRequest struct {
	RequestPacket
	TargetCharacterServerId int64 `json:",omitempty,omitzero"`
}
