package protos

type CafeSummonCharacterRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
	CharacterServerId int64 `json:",omitempty,omitzero"`
}
