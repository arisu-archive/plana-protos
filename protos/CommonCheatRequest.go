package protos

type CommonCheatRequest struct {
	RequestPacket
	Cheat                 string
	CharacterCustomPreset []*CheatCharacterCustomPreset
}
