package protos

type MomoTalkMessageListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CharacterDBId int64 `json:",omitempty,omitzero"`
}
