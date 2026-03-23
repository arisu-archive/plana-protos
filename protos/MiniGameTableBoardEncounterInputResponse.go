package protos

type MiniGameTableBoardEncounterInputResponse struct {
	ResponsePacket
	SaveDB                    *TBGBoardSaveDB `json:",omitempty,omitzero"`
	PlayerDiceResult          []int32
	PlayerAddDotEffectResult  *int32             `json:",omitempty,omitzero"`
	PlayerDicePlayResult      *TBGDiceRollResult `json:",omitempty,omitzero"`
	ParcelResultDB            *ParcelResultDB    `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB
}
