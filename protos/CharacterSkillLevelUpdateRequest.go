package protos

type CharacterSkillLevelUpdateRequest struct {
	RequestPacket
	TargetCharacterDBId int64                     `json:",omitempty,omitzero"`
	SkillSlot           SkillSlot                 `json:",omitempty,omitzero"`
	Level               int32                     `json:",omitempty,omitzero"`
	ReplaceInfos        []SelectTicketReplaceInfo `json:",omitempty,omitzero"`
}
