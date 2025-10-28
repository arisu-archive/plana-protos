package protos

type CraftRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	CraftInfos []CraftInfoDB `json:",omitempty,omitzero"`
}
