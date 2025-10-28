package protos

type CharacterSkillLevelUpdateRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetCharacterDBId int64 `json:",omitempty,omitzero"`
	SkillSlot SkillSlot `json:",omitempty,omitzero"`
	Level int32 `json:",omitempty,omitzero"`
	ReplaceInfos []SelectTicketReplaceInfo `json:",omitempty,omitzero"`
}
