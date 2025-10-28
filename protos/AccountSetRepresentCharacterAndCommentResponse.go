package protos

type AccountSetRepresentCharacterAndCommentResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountDB AccountDB `json:",omitempty,omitzero"`
	RepresentCharacterDB CharacterDB `json:",omitempty,omitzero"`
}
