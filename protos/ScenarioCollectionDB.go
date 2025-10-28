package protos

import (
	"time"
)

type ScenarioCollectionDB struct {
	GroupId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	ReceiveDate time.Time `json:",omitempty,omitzero"`
}
