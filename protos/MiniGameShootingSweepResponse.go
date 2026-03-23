package protos

type MiniGameShootingSweepResponse struct {
	ResponsePacket
	Rewards        [][]*ParcelInfo
	ParcelResultDB *ParcelResultDB `json:",omitempty,omitzero"`
}
