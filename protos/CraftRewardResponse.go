package protos

type CraftRewardResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	CraftInfos     []CraftInfoDB
}
