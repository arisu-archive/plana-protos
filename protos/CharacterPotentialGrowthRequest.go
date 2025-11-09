package protos

type CharacterPotentialGrowthRequest struct {
	RequestPacket
	TargetCharacterDBId       int64                      `json:",omitempty,omitzero"`
	PotentialGrowthRequestDBs []PotentialGrowthRequestDB `json:",omitempty,omitzero"`
	ReplaceInfos              []SelectTicketReplaceInfo  `json:",omitempty,omitzero"`
}
