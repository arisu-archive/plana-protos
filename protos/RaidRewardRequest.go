package protos

type RaidRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidServerId int64 `json:",omitempty,omitzero"`
	IsPractice bool `json:",omitempty,omitzero"`
}
