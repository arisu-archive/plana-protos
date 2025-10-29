package protos

type EventContentEnterStoryStageRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
}
