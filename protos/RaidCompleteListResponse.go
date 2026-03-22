package protos

type RaidCompleteListResponse struct {
	ResponsePacket
	RaidDBs           []RaidDB
	StackedDamage     int64 `json:",omitempty,omitzero"`
	ReceiveRewardId   []int64
	CurSeasonUniqueId int64 `json:",omitempty,omitzero"`
}
