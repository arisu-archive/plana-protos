package protos

type MiniGameCCGCreateGameRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	ForceDiscardSave bool `json:",omitempty,omitzero"`
	DisablePerk bool `json:",omitempty,omitzero"`
}
