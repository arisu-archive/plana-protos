package protos

type TacticalRelayHistoryDB struct {
	AccountId     int64  `json:",omitempty,omitzero"`
	SeasonId      int64  `json:",omitempty,omitzero"`
	StageId       int64  `json:",omitempty,omitzero"`
	BestClearWave int32  `json:",omitempty,omitzero"`
	LastClearDate MxTime `json:",omitempty,omitzero"`
}
