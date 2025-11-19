package protos

type EventContentStoryStageResultRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	StageUniqueId  int64 `json:",omitempty,omitzero"`
}
