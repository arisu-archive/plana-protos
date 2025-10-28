package protos

type MiniGameTableBoardMoveResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PlayerDB TBGPlayerDB `json:",omitempty,omitzero"`
	SaveDB TBGBoardSaveDB `json:",omitempty,omitzero"`
	EncounterDB TBGEncounterDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
