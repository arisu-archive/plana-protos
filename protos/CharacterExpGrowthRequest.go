package protos

type CharacterExpGrowthRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetCharacterServerId int64 `json:",omitempty,omitzero"`
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
}
