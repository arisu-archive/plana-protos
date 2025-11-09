package protos

type MiniGameTableBoardSweepResponse struct {
	ResponsePacket
	SaveDB         TBGBoardSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
