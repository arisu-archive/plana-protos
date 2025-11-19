package protos

type EventContentPurchasePlayCountHardStageRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	StageUniqueId  int64 `json:",omitempty,omitzero"`
}
