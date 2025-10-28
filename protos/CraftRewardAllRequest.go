package protos

type CraftRewardAllRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
