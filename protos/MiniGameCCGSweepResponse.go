package protos

type MiniGameCCGSweepResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Rewards [][]ParcelInfo `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
