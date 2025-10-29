package protos

type MiniGameDreamMakerEndingRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
