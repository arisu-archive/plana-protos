package protos

type MomoTalkMessageListRequest struct {
	RequestPacket
	CharacterDBId int64 `json:",omitempty,omitzero"`
}
