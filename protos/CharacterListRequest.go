package protos

type CharacterListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
