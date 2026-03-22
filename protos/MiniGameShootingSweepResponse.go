package protos

type MiniGameShootingSweepResponse struct {
	ResponsePacket
	Rewards        [][]ParcelInfo
	ParcelResultDB ParcelResultDB
}
