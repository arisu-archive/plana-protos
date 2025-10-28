package protos

type CharacterGearListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
