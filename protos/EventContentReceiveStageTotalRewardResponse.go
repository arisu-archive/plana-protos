package protos

type EventContentReceiveStageTotalRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	AlreadyReceiveRewardId []int64 `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
