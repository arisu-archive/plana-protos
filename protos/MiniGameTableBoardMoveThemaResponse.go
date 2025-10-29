package protos

type MiniGameTableBoardMoveThemaResponse struct {
	ResponsePacket
	SaveDB TBGBoardSaveDB `json:",omitempty,omitzero"`
}
