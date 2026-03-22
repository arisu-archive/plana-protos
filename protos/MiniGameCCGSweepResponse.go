package protos

type MiniGameCCGSweepResponse struct {
	ResponsePacket
	Rewards        [][]ParcelInfo
	ParcelResultDB ParcelResultDB
}
