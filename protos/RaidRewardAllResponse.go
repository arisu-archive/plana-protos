package protos

type RaidRewardAllResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
