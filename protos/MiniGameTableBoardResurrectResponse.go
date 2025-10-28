package protos

type MiniGameTableBoardResurrectResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PlayerDB TBGPlayerDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
