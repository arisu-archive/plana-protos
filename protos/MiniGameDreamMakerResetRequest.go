package protos

type MiniGameDreamMakerResetRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
