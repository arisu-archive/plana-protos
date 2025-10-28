package protos

type RaidRewardAllRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
