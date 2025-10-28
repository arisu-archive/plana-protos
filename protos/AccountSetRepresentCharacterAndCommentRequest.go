package protos

type AccountSetRepresentCharacterAndCommentRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RepresentCharacterServerId int64 `json:",omitempty,omitzero"`
	Comment string `json:",omitempty,omitzero"`
}
