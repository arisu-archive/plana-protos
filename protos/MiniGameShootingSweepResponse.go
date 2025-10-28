package protos

type MiniGameShootingSweepResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Rewards [][]ParcelInfo `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
