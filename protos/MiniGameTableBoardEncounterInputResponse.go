package protos

type MiniGameTableBoardEncounterInputResponse struct {
	ResponsePacket
	SaveDB                    TBGBoardSaveDB
	EncounterDB               TBGEncounterDB
	PlayerDiceResult          []int32
	PlayerAddDotEffectResult  *int32             `json:",omitempty,omitzero"`
	PlayerDicePlayResult      *TBGDiceRollResult `json:",omitempty,omitzero"`
	ParcelResultDB            ParcelResultDB
	EventContentCollectionDBs []EventContentCollectionDB
}
