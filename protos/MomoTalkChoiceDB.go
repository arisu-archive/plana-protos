package protos

type MomoTalkChoiceDB struct {
	CharacterDBId   int64  `json:",omitempty,omitzero"`
	MessageGroupId  int64  `json:",omitempty,omitzero"`
	ChosenMessageId int64  `json:",omitempty,omitzero"`
	ChosenDate      MxTime `json:",omitempty,omitzero"`
}
