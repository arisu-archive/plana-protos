package protos

type CafeSummonCharacterResponse struct {
	ResponsePacket
	CafeDB  CafeDB   `json:",omitempty,omitzero"`
	CafeDBs []CafeDB `json:",omitempty,omitzero"`
}
