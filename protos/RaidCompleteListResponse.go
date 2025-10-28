package protos

type RaidCompleteListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidDBs []RaidDB `json:",omitempty,omitzero"`
	StackedDamage int64 `json:",omitempty,omitzero"`
	ReceiveRewardId []int64 `json:",omitempty,omitzero"`
	CurSeasonUniqueId int64 `json:",omitempty,omitzero"`
}
