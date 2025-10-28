package protos

type MomoTalkReadRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CharacterDBId int64 `json:",omitempty,omitzero"`
	LastReadMessageGroupId int64 `json:",omitempty,omitzero"`
	ChosenMessageId *int64 `json:",omitempty,omitzero"`
}
