package protos

type WorldRaidClearHistoryDB struct {
	SeasonId          int64  `json:",omitempty,omitzero"`
	GroupId           int64  `json:",omitempty,omitzero"`
	RewardReceiveDate MxTime `json:",omitempty,omitzero"`
}
