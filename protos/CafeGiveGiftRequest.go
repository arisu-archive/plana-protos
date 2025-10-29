package protos

type CafeGiveGiftRequest struct {
	RequestPacket
	CafeDBId int64 `json:",omitempty,omitzero"`
	CharacterUniqueId int64 `json:",omitempty,omitzero"`
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
}
