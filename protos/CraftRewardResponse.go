package protos

type CraftRewardResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	CraftInfos []CraftInfoDB `json:",omitempty,omitzero"`
}
