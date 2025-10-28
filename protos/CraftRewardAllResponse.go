package protos

type CraftRewardAllResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	CraftInfos []CraftInfoDB `json:",omitempty,omitzero"`
}
