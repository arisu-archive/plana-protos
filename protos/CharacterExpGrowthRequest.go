package protos

type CharacterExpGrowthRequest struct {
	RequestPacket
	TargetCharacterServerId int64 `json:",omitempty,omitzero"`
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
}
