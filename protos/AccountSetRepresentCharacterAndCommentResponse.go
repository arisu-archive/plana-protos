package protos

type AccountSetRepresentCharacterAndCommentResponse struct {
	ResponsePacket
	AccountDB AccountDB `json:",omitempty,omitzero"`
	RepresentCharacterDB CharacterDB `json:",omitempty,omitzero"`
}
