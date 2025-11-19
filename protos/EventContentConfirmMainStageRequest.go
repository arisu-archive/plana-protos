package protos

type EventContentConfirmMainStageRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	StageUniqueId  int64 `json:",omitempty,omitzero"`
}
