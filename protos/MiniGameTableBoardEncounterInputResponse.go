package protos

type MiniGameTableBoardEncounterInputResponse struct {
	ResponsePacket
	SaveDB                    TBGBoardSaveDB
	PlayerDiceResult          []int32
	PlayerAddDotEffectResult  *int32             `json:",omitempty,omitzero"`
	PlayerDicePlayResult      *TBGDiceRollResult `json:",omitempty,omitzero"`
	ParcelResultDB            ParcelResultDB
	EventContentCollectionDBs []EventContentCollectionDB
}
