package protos

import (
	"time"
)

type ShopRecruitDB struct {
	Id int64 `json:",omitempty,omitzero"`
	SalesStartDate time.Time `json:",omitempty,omitzero"`
	SalesEndDate time.Time `json:",omitempty,omitzero"`
	UpdateDate time.Time `json:",omitempty,omitzero"`
}
