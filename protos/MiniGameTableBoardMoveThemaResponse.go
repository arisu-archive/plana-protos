package protos

type MiniGameTableBoardMoveThemaResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDB TBGBoardSaveDB `json:",omitempty,omitzero"`
}
