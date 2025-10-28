package protos

import (
	"time"
)

type PickupFirstGetHistoryDB struct {
	ShopRecruitId int64 `json:",omitempty,omitzero"`
	LogDate time.Time `json:",omitempty,omitzero"`
}
