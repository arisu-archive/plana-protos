package protos

import (
	"time"
)

type EventContentCollectionDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	GroupId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	ReceiveDate time.Time `json:",omitempty,omitzero"`
}
