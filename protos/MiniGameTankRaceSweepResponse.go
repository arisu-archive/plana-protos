package protos

type MiniGameTankRaceSweepResponse struct {
	ResponsePacket
	ParcelResultDB     *ParcelResultDB `json:",omitempty,omitzero"`
	RewardBySweepCount [][]*ParcelInfo
}
