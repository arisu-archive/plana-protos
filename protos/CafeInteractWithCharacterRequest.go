package protos

type CafeInteractWithCharacterRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
	CharacterId int64 `json:",omitempty,omitzero"`
}
