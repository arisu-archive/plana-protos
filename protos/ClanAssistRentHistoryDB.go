package protos

type ClanAssistRentHistoryDB struct {
	AssistCharacterAccountId int64  `json:",omitempty,omitzero"`
	AssistCharacterDBId      int64  `json:",omitempty,omitzero"`
	RentDate                 MxTime `json:",omitempty,omitzero"`
	AssistCharacterId        int64  `json:",omitempty,omitzero"`
}
