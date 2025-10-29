package protos

type MiniGameCCGRerollRewardRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
