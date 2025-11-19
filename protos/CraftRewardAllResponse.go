package protos

type CraftRewardAllResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	CraftInfos     []CraftInfoDB  `json:",omitempty,omitzero"`
}
