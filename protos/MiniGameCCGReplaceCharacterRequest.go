package protos

type MiniGameCCGReplaceCharacterRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	SlotIndex int32 `json:",omitempty,omitzero"`
	CharacterId int64 `json:",omitempty,omitzero"`
	IsStriker bool `json:",omitempty,omitzero"`
}
