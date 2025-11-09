package protos

type MomoTalkOutLineDB struct {
	CharacterDBId        int64  `json:",omitempty,omitzero"`
	CharacterId          int64  `json:",omitempty,omitzero"`
	LatestMessageGroupId int64  `json:",omitempty,omitzero"`
	ChosenMessageId      *int64 `json:",omitempty,omitzero"`
	LastUpdateDate       MxTime `json:",omitempty,omitzero"`
}
