package protos

type CraftRewardAllResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	CraftInfos     []CraftInfoDB
}
