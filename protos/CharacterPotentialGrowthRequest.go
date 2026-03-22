package protos

type CharacterPotentialGrowthRequest struct {
	RequestPacket
	TargetCharacterDBId       int64 `json:",omitempty,omitzero"`
	PotentialGrowthRequestDBs []PotentialGrowthRequestDB
	ReplaceInfos              []SelectTicketReplaceInfo
}
