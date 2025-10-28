package protos

import (
	"time"
)

type ShopFreeRecruitHistoryDB struct {
	UniqueId int64 `json:",omitempty,omitzero"`
	RecruitCount int32 `json:",omitempty,omitzero"`
	LastUpdateDate time.Time `json:",omitempty,omitzero"`
}
