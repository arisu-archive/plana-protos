package protos

import (
	"time"
)

type EmblemDB struct {
	ParcelBase
	UniqueId int64 `json:",omitempty,omitzero"`
	ReceiveDate time.Time `json:",omitempty,omitzero"`
}
