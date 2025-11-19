package protos

type EventContentEnterSubStageRequest struct {
	RequestPacket
	EventContentId              int64 `json:",omitempty,omitzero"`
	StageUniqueId               int64 `json:",omitempty,omitzero"`
	LastEnterStageEchelonNumber int64 `json:",omitempty,omitzero"`
}
