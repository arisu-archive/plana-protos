package protos

import (
	"time"
)

type MomoTalkChoiceDB struct {
	CharacterDBId int64 `json:",omitempty,omitzero"`
	MessageGroupId int64 `json:",omitempty,omitzero"`
	ChosenMessageId int64 `json:",omitempty,omitzero"`
	ChosenDate time.Time `json:",omitempty,omitzero"`
}
