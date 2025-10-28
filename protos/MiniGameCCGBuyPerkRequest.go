package protos

type MiniGameCCGBuyPerkRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	PerkId int64 `json:",omitempty,omitzero"`
}
