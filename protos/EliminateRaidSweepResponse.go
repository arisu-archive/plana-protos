package protos

type EliminateRaidSweepResponse struct {
	ResponsePacket
	TotalSeasonPoint int64 `json:",omitempty,omitzero"`
	Rewards          [][]ParcelInfo
	ParcelResultDB   *ParcelResultDB `json:",omitempty,omitzero"`
}
