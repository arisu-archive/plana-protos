package protos

type CafeInteractWithCharacterRequest struct {
	RequestPacket
	CafeDBId int64 `json:",omitempty,omitzero"`
	CharacterId int64 `json:",omitempty,omitzero"`
}
