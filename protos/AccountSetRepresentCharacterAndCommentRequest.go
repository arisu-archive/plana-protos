package protos

type AccountSetRepresentCharacterAndCommentRequest struct {
	RequestPacket
	RepresentCharacterServerId int64 `json:",omitempty,omitzero"`
	Comment string `json:",omitempty,omitzero"`
}
