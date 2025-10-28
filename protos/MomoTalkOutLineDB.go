package protos

import (
	"time"
)

type MomoTalkOutLineDB struct {
	CharacterDBId int64 `json:",omitempty,omitzero"`
	CharacterId int64 `json:",omitempty,omitzero"`
	LatestMessageGroupId int64 `json:",omitempty,omitzero"`
	ChosenMessageId *int64 `json:",omitempty,omitzero"`
	LastUpdateDate time.Time `json:",omitempty,omitzero"`
}
