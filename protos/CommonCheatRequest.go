package protos

type CommonCheatRequest struct {
	RequestPacket
	Cheat string `json:",omitempty,omitzero"`
	CharacterCustomPreset []CheatCharacterCustomPreset `json:",omitempty,omitzero"`
}
