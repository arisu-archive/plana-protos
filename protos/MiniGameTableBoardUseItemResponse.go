package protos

type MiniGameTableBoardUseItemResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PlayerDB TBGPlayerDB `json:",omitempty,omitzero"`
}
