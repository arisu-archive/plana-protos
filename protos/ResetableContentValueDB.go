package protos

import (
	"time"
)

type ResetableContentValueDB struct {
	ResetableContentId ResetableContentId `json:",omitempty,omitzero"`
	ContentValue int64 `json:",omitempty,omitzero"`
	LastUpdateTime time.Time `json:",omitempty,omitzero"`
}
