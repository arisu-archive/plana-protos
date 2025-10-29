package protos

type EventContentReceiveStageTotalRewardRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
