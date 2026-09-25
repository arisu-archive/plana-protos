package protos

type EventContentClueSearchBatchSubmitRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	SlotNumbers    []int64
}
