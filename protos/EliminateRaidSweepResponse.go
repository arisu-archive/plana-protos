package protos

type EliminateRaidSweepResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TotalSeasonPoint int64 `json:",omitempty,omitzero"`
	Rewards [][]ParcelInfo `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
