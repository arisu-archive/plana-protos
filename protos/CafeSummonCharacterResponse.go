package protos

type CafeSummonCharacterResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDB CafeDB `json:",omitempty,omitzero"`
	CafeDBs []CafeDB `json:",omitempty,omitzero"`
}
