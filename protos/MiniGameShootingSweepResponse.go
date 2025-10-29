package protos

type MiniGameShootingSweepResponse struct {
	ResponsePacket
	Rewards [][]ParcelInfo `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
