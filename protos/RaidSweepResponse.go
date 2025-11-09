package protos

type RaidSweepResponse struct {
	ResponsePacket
	TotalSeasonPoint int64          `json:",omitempty,omitzero"`
	Rewards          [][]ParcelInfo `json:",omitempty,omitzero"`
	ParcelResultDB   ParcelResultDB `json:",omitempty,omitzero"`
}
