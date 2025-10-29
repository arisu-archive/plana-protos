package protos

type MiniGameTableBoardSyncResponse struct {
	ResponsePacket
	SaveDB TBGBoardSaveDB `json:",omitempty,omitzero"`
}
