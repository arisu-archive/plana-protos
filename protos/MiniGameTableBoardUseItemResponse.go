package protos

type MiniGameTableBoardUseItemResponse struct {
	ResponsePacket
	PlayerDB TBGPlayerDB `json:",omitempty,omitzero"`
}
