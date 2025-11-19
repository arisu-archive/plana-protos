package protos

type MiniGameTableBoardClearThemaResponse struct {
	ResponsePacket
	SaveDB         TBGBoardSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
