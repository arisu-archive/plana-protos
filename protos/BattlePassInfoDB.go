package protos

import (
	"time"
)

type BattlePassInfoDB struct {
	BattlePassId int64 `json:",omitempty,omitzero"`
	PassLevel int32 `json:",omitempty,omitzero"`
	PassExp int64 `json:",omitempty,omitzero"`
	PurchaseGroupId int64 `json:",omitempty,omitzero"`
	ReceiveRewardLevel int32 `json:",omitempty,omitzero"`
	ReceivePurchaseRewardLevel int32 `json:",omitempty,omitzero"`
	WeeklyPassExp int64 `json:",omitempty,omitzero"`
	LastWeeklyPassExpLimitRefreshDate time.Time `json:",omitempty,omitzero"`
}
