package protos

type MiniGameTableBoardSyncResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDB TBGBoardSaveDB `json:",omitempty,omitzero"`
}
