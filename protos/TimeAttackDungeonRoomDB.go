package protos

type TimeAttackDungeonRoomDB struct {
	AccountId         int64                              `json:",omitempty,omitzero"`
	SeasonId          int64                              `json:",omitempty,omitzero"`
	RoomId            int64                              `json:",omitempty,omitzero"`
	CreateDate        MxTime                             `json:",omitempty,omitzero"`
	RewardDate        MxTime                             `json:",omitempty,omitzero"`
	IsPractice        bool                               `json:",omitempty,omitzero"`
	SweepHistoryDates []MxTime                           `json:",omitempty,omitzero"`
	BattleHistoryDBs  []TimeAttackDungeonBattleHistoryDB `json:",omitempty,omitzero"`
}
