package protos

import (
	"time"
)

type TimeAttackDungeonRoomDB struct {
	AccountId int64 `json:",omitempty,omitzero"`
	SeasonId int64 `json:",omitempty,omitzero"`
	RoomId int64 `json:",omitempty,omitzero"`
	CreateDate time.Time `json:",omitempty,omitzero"`
	RewardDate time.Time `json:",omitempty,omitzero"`
	IsPractice bool `json:",omitempty,omitzero"`
	SweepHistoryDates []time.Time `json:",omitempty,omitzero"`
	BattleHistoryDBs []TimeAttackDungeonBattleHistoryDB `json:",omitempty,omitzero"`
}
