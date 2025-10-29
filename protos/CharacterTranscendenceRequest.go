package protos

type CharacterTranscendenceRequest struct {
	RequestPacket
	TargetCharacterServerId int64 `json:",omitempty,omitzero"`
}
