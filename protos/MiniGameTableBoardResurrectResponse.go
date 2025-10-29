package protos

type MiniGameTableBoardResurrectResponse struct {
	ResponsePacket
	PlayerDB TBGPlayerDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
