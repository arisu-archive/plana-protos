package protos

type RaidRewardRequest struct {
	RequestPacket
	RaidServerId int64 `json:",omitempty,omitzero"`
	IsPractice   bool  `json:",omitempty,omitzero"`
}
