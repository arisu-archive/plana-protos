package protos

type MiniGameDreamMakerDailyClosingRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
