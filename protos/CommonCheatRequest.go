package protos

type CommonCheatRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Cheat string `json:",omitempty,omitzero"`
	CharacterCustomPreset []CheatCharacterCustomPreset `json:",omitempty,omitzero"`
}
