package protos

type MiniGameCCGBuyPerkRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	PerkId         int64 `json:",omitempty,omitzero"`
}
