package protos

import (
	"time"
)

type ClanAssistRentHistoryDB struct {
	AssistCharacterAccountId int64 `json:",omitempty,omitzero"`
	AssistCharacterDBId int64 `json:",omitempty,omitzero"`
	RentDate time.Time `json:",omitempty,omitzero"`
	AssistCharacterId int64 `json:",omitempty,omitzero"`
}
