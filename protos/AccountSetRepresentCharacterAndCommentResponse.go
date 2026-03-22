package protos

type AccountSetRepresentCharacterAndCommentResponse struct {
	ResponsePacket
	AccountDB            AccountDB
	RepresentCharacterDB CharacterDB
}
