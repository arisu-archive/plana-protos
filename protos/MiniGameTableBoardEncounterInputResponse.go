package protos

type MiniGameTableBoardEncounterInputResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDB TBGBoardSaveDB `json:",omitempty,omitzero"`
	EncounterDB TBGEncounterDB `json:",omitempty,omitzero"`
	PlayerDiceResult []int32 `json:",omitempty,omitzero"`
	PlayerAddDotEffectResult *int32 `json:",omitempty,omitzero"`
	PlayerDicePlayResult *TBGDiceRollResult `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
