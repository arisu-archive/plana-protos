package protos

type EventContentEnterStoryStageRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
}
