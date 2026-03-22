package protos

type EventContentReceiveStageTotalRewardResponse struct {
	ResponsePacket
	EventContentId         int64 `json:",omitempty,omitzero"`
	AlreadyReceiveRewardId []int64
	ParcelResultDB         ParcelResultDB
}
